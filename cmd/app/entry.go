package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/haski007/insta-bot/pkg/graceful"
	"github.com/haski007/insta-bot/pkg/run"
	"github.com/haski007/pretty"
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

	logrus.Println("Run func")
	logrus.Println("cfg:", pretty.String(cfg))

	httpMux := http.NewServeMux()
	httpMux.Handle("/metrics", promhttp.Handler())

	metricsServer := &http.Server{Addr: args.MetricsAddr, Handler: httpMux}

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramBot.Token)
	if err != nil {
		return fmt.Errorf("new tg bot api err: %w", err)
	}

	var server errgroup.Group

	server.Go(func() error {
		defer stop()

		time.Sleep(time.Second * 4)
		return nil
	})

	server.Go(func() error {
		defer stop()
		log.Infof("metrics service listening on %s", args.MetricsAddr)

		if errLA := metricsServer.ListenAndServe(); errLA != nil && !errors.Is(errLA, http.ErrServerClosed) {
			logrus.Errorf("metrics server exit with error: %s", errLA)
		}

		return nil
	})

	go graceful.Shutdown(
		ctx,
		graceful.TGBOT(bot, cfg.TelegramBot.CreatorUserID),
		graceful.HTTP(metricsServer),
		graceful.CloseFunc(func() error {
			stop()
			return nil
		}),
	)

	return server.Wait()
}
