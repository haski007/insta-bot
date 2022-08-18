package listener

import (
	"github.com/haski007/insta-bot/pkg/emoji"
	"github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (rcv *InstaBotService) cmdTestHandler(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	if err := rcv.SendMessage(chatID, "Pososi! "+emoji.Gear); err != nil {
		logrus.WithError(err).Println("[cmdTestHandler] send message to chat")
	}
}
