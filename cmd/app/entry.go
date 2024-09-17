package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/go-redis/redis"
	"github.com/haski007/insta-bot/internal/bot/listener"
	"github.com/haski007/insta-bot/internal/clients/chatgpt"
	"github.com/haski007/insta-bot/internal/clients/instapi"
	"github.com/haski007/insta-bot/internal/clients/tiktokapi"
	"github.com/haski007/insta-bot/internal/clients/youtube"
	"github.com/haski007/insta-bot/pkg/graceful"
	"github.com/haski007/insta-bot/pkg/run"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sashabaranov/go-openai"
	"github.com/sethvargo/go-envconfig"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	redisWrapper "github.com/haski007/insta-bot/internal/storage/redis"
)

func Run(ctx context.Context, args run.Args) error {
	log := logrus.New()
	log.SetLevel(args.LogLevel)
	log.SetFormatter(&logrus.JSONFormatter{})

	var cfg Config
	if err := envconfig.Process(ctx, &cfg); err != nil {
		return fmt.Errorf("load config from env err: %w", err)
	}

	// ---> Google AUTH
	//b := bytes.NewBufferString(cfg.Clients.Google.Credentials).Bytes()

	// If modifying these scopes, delete your previously saved token.json.
	//config, err := google.ConfigFromJSON(b,
	//	calendar.CalendarScope,
	//	calendar.CalendarReadonlyScope,
	//	calendar.CalendarEventsScope)
	//if err != nil {
	//	log.Fatalf("Unable to parse client secret file to config: %v", err)
	//}
	//gClient, gTokenSource := googleWrapper.GetClient(config)

	//gSrv, err := calendar.NewService(ctx, option.WithHTTPClient(gClient))
	//if err != nil {
	//	log.Fatalf("Unable to retrieve Calendar client: %v", err)
	//}
	//
	//calendarSrv := calendarWrapper.New(gSrv, gTokenSource, config)

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)

	httpMux := http.NewServeMux()
	httpMux.Handle("/metrics", promhttp.Handler())

	metricsServer := &http.Server{Addr: args.MetricsAddr, Handler: httpMux}

	botApi, err := tgbotapi.NewBotAPI(cfg.TelegramBot.Token)
	if err != nil {
		return fmt.Errorf("new tg bot api token: [%s] err: %w", cfg.TelegramBot.Token, err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = int(cfg.TelegramBot.UpdatesTimeout.Seconds())
	chUpdates := botApi.GetUpdatesChan(u)

	apiCli := instapi.New()

	redCC := redis.NewClient(&redis.Options{
		Addr:     cfg.Clients.Redis.Addr,
		Password: cfg.Clients.Redis.Pass,
	})
	defer redCC.Close()

	redisStorage, err := redisWrapper.NewClient(
		redCC,
		cfg.Clients.Redis.ConversationTTL,
		cfg.Clients.Redis.HistoryMessagesTTL,
	)
	if err != nil {
		return fmt.Errorf("connect to redis err: %w", err)
	}

	// ---> open AI
	ai := openai.NewClient(cfg.Clients.OpenAI.ApiKey)
	chatGptSrv, err := chatgpt.NewService(
		ai,
		cfg.Clients.OpenAI.GPTModelForConv,
		cfg.Clients.OpenAI.GPTModelForHistory,
		cfg.Clients.OpenAI.ApiKey,
	)
	if err != nil {
		return fmt.Errorf("chat gpt service err: %w", err)
	}

	botSrv := listener.NewInstaBotService(
		ctx,
		botApi,
		apiCli,
		cfg.TelegramBot.CreatorUserID,
		chUpdates,
		cfg.CaptionCharsLimit,
		tiktokapi.New(),
		youtube.New(cfg.Clients.YoutubeApi.MaxQuality),
		redisStorage,
		nil,
		chatGptSrv,
	).SetLogger(log)

	// reads from redis all the funcs that should be run in set time
	if err := botSrv.RunAfterFuncsPolls(); err != nil {
		return fmt.Errorf("run afterFunc polls err: %w", err)
	}

	// run a monitor that checks if redis is not read only
	go botSrv.RedisMonitor()

	if err := tgbotapi.SetLogger(log); err != nil {
		return fmt.Errorf("set looger for tgbotapi package err: %w", err)
	}

	var server errgroup.Group

	server.Go(func() error {
		defer stop()

		me, er := botApi.GetMe()
		if er != nil {
			logrus.WithError(err).Println("bot api getMe")
		}

		log.Infof("bot @%s is polling now", me.UserName)

		if errL := botSrv.StartPool(); errL != nil {
			logrus.WithError(err).Println("bot listener exit with error")
		}

		return nil
	})

	server.Go(func() error {
		defer stop()
		log.Infof("metrics service listening on %s", args.MetricsAddr)

		if errLA := metricsServer.ListenAndServe(); errLA != nil && !errors.Is(errLA, http.ErrServerClosed) {
			logrus.WithError(err).Println("metrics server exit with error")
		}

		return nil
	})

	// Run newsletter monitor
	server.Go(func() error {
		defer stop()
		log.Infof("Newsletter monitor is running")

		botSrv.RunNewsLetter()
		return nil
	})

	go graceful.Shutdown(
		ctx,
		graceful.TGBOT(botSrv),
		graceful.HTTP(metricsServer),
		graceful.CloseFunc(func() error {
			stop()
			return nil
		}),
	)

	return server.Wait()
}
