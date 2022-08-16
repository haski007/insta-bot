package graceful

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/haski007/insta-bot/pkg/factory"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/sirupsen/logrus"
)

const shutdownTimeout = 10 * time.Second

type CloseFunc func() error

func (cf CloseFunc) Close() error { return cf() }

func CloseAll(cc ...io.Closer) CloseFunc {
	return func() error {
		for _, c := range cc {
			if err := c.Close(); err != nil {
				logrus.Errorf("graceful shutdown: %v", err)
			}
		}

		return nil
	}
}

func HTTP(srv *http.Server) CloseFunc {
	return func() error {
		ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return nil
			}

			return err
		}

		return nil
	}
}

func TGBOT(srv *tgbotapi.BotAPI, creator int64) CloseFunc {
	return func() error {
		if err := factory.NotifyCreator(
			srv,
			creator,
			"Bot got signal so it's shutting down...",
		); err != nil {
			return fmt.Errorf("gracefull shutdown err: %w", err)
		}
		return nil
	}
}

func Shutdown(ctx context.Context, cc ...io.Closer) {
	<-ctx.Done()

	logrus.Infof("received interrupting signal, terminating...")

	if err := CloseAll(cc...).Close(); err != nil {
		logrus.Errorf("graceful shutdown: %v", err)
	}
}
