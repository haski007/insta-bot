package listener

import (
	"github.com/haski007/insta-bot/pkg/emoji"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (rcv *InstaBotService) cmdPurgeHistory(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	if err := rcv.storage.PurgeHistory(chatID); err != nil {
		rcv.SendError(chatID, ErrInternalServerError)
		rcv.log.WithError(err).Error("[cmdPurgeHistory] purge chat history")
		return
	}

	if _, err := rcv.bot.Send(tgbotapi.NewMessage(chatID, "History purged "+emoji.Basket)); err != nil {
		rcv.SendError(chatID, ErrInternalServerError)
		rcv.log.WithError(err).Error("[cmdPurgeHistory] chat history purged")
		return
	}
}
