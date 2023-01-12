package listener

import (
	"github.com/haski007/insta-bot/pkg/emoji"
	"github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (rcv *InstaBotService) cmdPurgeCSGOPlayersHandler(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	if err := rcv.storage.DeleteChat(CSGOContext, chatID); err != nil {
		rcv.SendError(chatID, ErrInternalServerError)
		rcv.log.WithError(err).Error("[cmdPurgeCSGOPlayersHandler] delete csgo chat with members")
		return
	}

	message := "Successfully deleted stopped to listen cs go triggers "
	if err := rcv.SendMessage(chatID, message+emoji.Basket); err != nil {
		logrus.WithError(err).Println("[cmdPurgeCSGOPlayersHandler] send message to chat")
	}
}
