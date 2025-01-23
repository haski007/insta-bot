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
	rcv.NotifyCreator(fmt.Sprintf("Bot was added to a new chat %d, chat title: %s", update.MyChatMember.Chat.ID, update.MyChatMember.Chat.Title))
	return rcv.SendMessage(update.MyChatMember.Chat.ID, message)
}
