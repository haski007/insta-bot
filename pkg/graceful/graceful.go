package graceful

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/haski007/insta-bot/internal/bot"

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

func TGBOT(srv bot.TgBot) CloseFunc {
	return func() error {
		ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		if err := srv.StopPool(ctx); err != nil {
			return err
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
