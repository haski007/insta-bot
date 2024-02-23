package listener

import (
	"github.com/haski007/insta-bot/pkg/emoji"
	"github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (rcv *InstaBotService) cmdPurgeFinalsPlayersHandler(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	if err := rcv.storage.DeleteChat(FinalsContext, chatID); err != nil {
		rcv.SendError(chatID, ErrInternalServerError)
		rcv.log.WithError(err).Error("[cmdPurgeFinalsPlayersHandler] delete finals chat with members")
		return
	}

	message := "Successfully deleted stopped to listen The Finals triggers "
	if err := rcv.SendMessage(chatID, message+emoji.Basket); err != nil {
		logrus.WithError(err).Println("[cmdPurgeFinalsPlayersHandler] send message to chat")
	}
}
