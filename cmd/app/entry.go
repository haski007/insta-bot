package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-redis/redis"
	"github.com/haski007/insta-bot/internal/bot/listener"
	arcraiders "github.com/haski007/insta-bot/internal/clients/arc-raiders"
	"github.com/haski007/insta-bot/internal/clients/chatgpt"
	"github.com/haski007/insta-bot/internal/clients/instloader"
	"github.com/haski007/insta-bot/internal/clients/tiktokapi"
	"github.com/haski007/insta-bot/internal/clients/youtube"
	"github.com/haski007/insta-bot/pkg/graceful"
	"github.com/haski007/insta-bot/pkg/run"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sashabaranov/go-openai"
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
	if err := Load(&cfg); err != nil {
		return fmt.Errorf("load config err: %w", err)
	}

	// ---> Google AUTH
	// b, err := os.ReadFile(cfg.Clients.Google.CredentialsPath)
	// if err != nil {
	// 	log.Fatalf("Unable to read client secret file: %v", err)
	// }

	// If modifying these scopes, delete your previously saved token.json.
	// config, err := google.ConfigFromJSON(b,
	// 	calendar.CalendarScope,
	// 	calendar.CalendarReadonlyScope,
	// 	calendar.CalendarEventsScope)
	// if err != nil {
	// 	log.Fatalf("Unable to parse client secret file to config: %v", err)
	// }
	// gClient, gTokenSource := googleWrapper.GetClient(config)

	// gSrv, err := calendar.NewService(ctx, option.WithHTTPClient(gClient))
	// if err != nil {
	// 	log.Fatalf("Unable to retrieve Calendar client: %v", err)
	// }

	// calendarSrv := calendarWrapper.New(gSrv, gTokenSource, config)

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)

	httpMux := http.NewServeMux()
	httpMux.Handle("/metrics", promhttp.Handler())

	metricsServer := &http.Server{Addr: args.MetricsAddr, Handler: httpMux}

	botApi, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		return fmt.Errorf("new tg bot api token: [%s] err: %w", cfg.Token, err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = cfg.UpdatesTimeoutSec
	chUpdates := botApi.GetUpdatesChan(u)

	// Create instloader client
	instloaderURL, err := url.Parse(cfg.InstloaderBaseURL)
	if err != nil {
		return fmt.Errorf("parse instloader URL err: %w", err)
	}
	instloaderClient := instloader.NewClient(instloaderURL)

	redCC := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPass,
	})
	defer redCC.Close()

	redisStorage, err := redisWrapper.NewClient(
		redCC,
		time.Minute*cfg.RedisConversationTTL,
		time.Hour*cfg.RedisHistoryMessagesTTL,
	)
	if err != nil {
		return fmt.Errorf("connect to redis err: %w", err)
	}

	// ---> open AI
	ai := openai.NewClient(cfg.OpenAIAPIKey)
	chatGptSrv, err := chatgpt.NewService(
		ai,
		cfg.OpenAIGPTModelForConv,
		cfg.OpenAIGPTModelForHistory,
		cfg.OpenAIAPIKey,
	)
	if err != nil {
		return fmt.Errorf("chat gpt service err: %w", err)
	}

	// ---> ARC raiders
	arcRaidersURL, err := url.Parse(cfg.ARCRAidersBaseURL)
	if err != nil {
		return fmt.Errorf("parse ARC raiders URL err: %w", err)
	}
	arcRaidersClient := arcraiders.NewClient(arcRaidersURL)
	if err != nil {
		return fmt.Errorf("new ARC raiders client err: %w", err)
	}


	botSrv := listener.NewInstaBotService(
		ctx,
		botApi,
		instloaderClient,
		cfg.CreatorUserID,
		chUpdates,
		cfg.CaptionCharsLimit,
		tiktokapi.New(),
		youtube.New(cfg.MaxQuality),
		redisStorage,
		nil,
		chatGptSrv,
		arcRaidersClient,
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

	server.Go(func() error {
		defer stop()
		log.Infof("ARC raiders monitor is running")

		botSrv.RunARCMonitor()
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
