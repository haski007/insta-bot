package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/haski007/insta-bot/internal/bot/publisher"
	"github.com/haski007/insta-bot/internal/clients/instapi"

	"github.com/haski007/insta-bot/internal/bot/listener"
	"github.com/haski007/insta-bot/pkg/graceful"
	"github.com/haski007/insta-bot/pkg/run"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Run(ctx context.Context, args run.Args) error {
	log := logrus.New()
	log.SetLevel(args.LogLevel)
	log.SetFormatter(&logrus.JSONFormatter{})

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)

	var cfg Config
	if err := Load(args.ConfigFile, &cfg); err != nil {
		return fmt.Errorf("load config %s err: %w", args.ConfigFile, err)
	}

	httpMux := http.NewServeMux()
	httpMux.Handle("/metrics", promhttp.Handler())

	metricsServer := &http.Server{Addr: args.MetricsAddr, Handler: httpMux}

	botApi, err := tgbotapi.NewBotAPI(cfg.TelegramBot.Token)
	if err != nil {
		return fmt.Errorf("new tg bot api err: %w", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = cfg.TelegramBot.UpdatesTimeoutSec
	chUpdates := botApi.GetUpdatesChan(u)

	instapiCfg := instapi.NewConfiguration()
	instapiCfg.Host = cfg.Clients.Instapi.Addr
	instapiCfg.Scheme = "http"
	instapiCli := instapi.NewAPIClient(instapiCfg)

	apiCli := publisher.New(
		instapiCli,
		cfg.Clients.Instapi.Username,
		cfg.Clients.Instapi.Password)

	// TODO: uncomment after debug
	//if err := apiCli.Login(ctx); err != nil {
	//	return fmt.Errorf("instapi entry login err: %w", err)
	//}

	botSrv := listener.NewInstaBotService(
		ctx,
		botApi,
		apiCli,
		cfg.TelegramBot.CreatorUserID,
		chUpdates,
		cfg.CaptionCharsLimit,
	).SetLogger(log)

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
