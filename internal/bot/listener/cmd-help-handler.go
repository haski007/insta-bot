package listener

import (
	"github.com/haski007/insta-bot/pkg/emoji"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (rcv *InstaBotService) cmdStartHandler(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	if err := rcv.sendStartInfo(chatID); err != nil {
		rcv.log.WithError(err).Println("[cmdStartHandler] send start info")
	}
}

func (rcv *InstaBotService) sendStartInfo(chatID int64) error {
	message := `Hello, I am a bot that can help you make your chat meme life a lot easier and funnier.
For the best user experience give me please administrator rights ` + emoji.FaceWinking
	return rcv.SendMessage(chatID, message)
}
