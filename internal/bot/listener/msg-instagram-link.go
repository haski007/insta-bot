package listener

import (
	"github.com/haski007/insta-bot/pkg/emoji"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (rcv *InstaBotService) msgInstagramLinkTrigger(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	messageID := update.Message.MessageID

	if err := rcv.DeleteMessage(chatID, messageID); err != nil {
		rcv.log.WithError(err).Error("[msgInstagramLinkTrigger] delete message")
	}

	if err := rcv.SendMessage(chatID, "I am gay! "+emoji.Check); err != nil {
		rcv.log.WithError(err).Error("[msgInstagramLinkTrigger] send message")
	}
}
