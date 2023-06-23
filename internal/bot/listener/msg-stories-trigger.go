package listener

import (
	"github.com/haski007/insta-bot/pkg/emoji"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (rcv *InstaBotService) msgStoriesTrigger(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	rcv.SendError(chatID, "Sorry, but this feature is not available yet "+emoji.SadFace)
}
