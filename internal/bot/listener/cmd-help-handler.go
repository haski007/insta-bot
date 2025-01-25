package listener

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (rcv *InstaBotService) cmdStartHandler(update tgbotapi.Update) {
	if err := rcv.sendStartInfo(update); err != nil {
		rcv.log.WithError(err).Println("[cmdStartHandler] send start info")
	}
}

func (rcv *InstaBotService) sendStartInfo(update tgbotapi.Update) error {
	message := `Нащо ви мене знову розбудили ?! 😡`

	var chatID int64
	var chatTitle string
	if update.MyChatMember != nil {
		chatID = update.MyChatMember.Chat.ID
		chatTitle = update.MyChatMember.Chat.Title
	} else if update.Message != nil {
		chatID = update.Message.Chat.ID
		chatTitle = update.Message.Chat.Title
	}
	rcv.NotifyCreator(fmt.Sprintf("Bot was added to a new chat %d, chat title: %s", chatID, chatTitle))
	return rcv.SendMessage(chatID, message)
}
