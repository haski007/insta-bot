package listener

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (rcv *InstaBotService) cmdStartHandler(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	if err := rcv.sendStartInfo(chatID); err != nil {
		rcv.log.WithError(err).Println("[cmdStartHandler] send start info")
	}
}

func (rcv *InstaBotService) sendStartInfo(chatID int64) error {
	message := `Нащо ви мене знову розбудили ?! 😡`
	return rcv.SendMessage(chatID, message)
}
