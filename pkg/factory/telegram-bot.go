package factory

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBotCfg struct {
	Token         string `json:"token" yaml:"token"`
	CreatorUserID int64  `json:"creator_user_id" yaml:"creator_user_id"`
}

func NotifyCreator(bot *tgbotapi.BotAPI, creator int64, message string) error {
	msg := tgbotapi.NewMessage(creator, message)
	msg.ParseMode = tgbotapi.ModeMarkdown
	_, err := bot.Send(msg)
	if err != nil {
		return fmt.Errorf("notify creator err: %w", err)
	}
	return nil
}
