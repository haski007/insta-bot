package listener

import (
	"github.com/haski007/insta-bot/pkg/emoji"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (rcv *InstaBotService) SendError(chatID int64, msg string) error {
	answer := tgbotapi.NewMessage(chatID, emoji.Warning+msg+emoji.Warning)

	if _, errN := rcv.bot.Send(answer); errN != nil {
		return errN
	}

	return nil
}
