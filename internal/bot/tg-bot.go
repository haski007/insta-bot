package bot

import (
	"context"
)

type TgBot interface {
	SendMessage(chatID int64, message string) error
	DeleteMessage(chatID int64, messageID int) error

	StartPool() error
	StopPool(ctx context.Context) error
}
