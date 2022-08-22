package bot

import (
	"context"
	"errors"
)

type TgBot interface {
	SendMessage(chatID int64, message string) error
	SendMessageWithoutMarkdown(chatID int64, message string) error
	DeleteMessage(chatID int64, messageID int) error
	NotifyCreator(message string) error
	SendError(chatID int64, msg string) error

	StartPool() error
	StopPool(ctx context.Context) error
}

var (
	ErrWrongFileFormat = errors.New("WRONG_FILE_FORMAT")
)
