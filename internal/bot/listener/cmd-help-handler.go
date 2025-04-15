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
	if rcv == nil {
		return fmt.Errorf("InstaBotService is nil")
	}

	message := `Нащо ви мене знову розбудили ?! 😡`

	var chatID int64
	var chatTitle string

	// Try to get chat info from MyChatMember first
	if update.MyChatMember != nil {
		chatID = update.MyChatMember.Chat.ID
		chatTitle = update.MyChatMember.Chat.Title
	} else if update.Message != nil {
		// Fallback to Message if MyChatMember is not available
		chatID = update.Message.Chat.ID
		chatTitle = update.Message.Chat.Title
	} else {
		return fmt.Errorf("no valid chat information found in update")
	}

	// Only notify creator if we have valid chat information
	if chatID != 0 {
		if err := rcv.NotifyCreator(fmt.Sprintf("Bot was added to a new chat %d, chat title: %s", chatID, chatTitle)); err != nil {
			// Log the error but don't return it since it's not critical
			if rcv.log != nil {
				rcv.log.WithError(err).Error("[sendStartInfo] failed to notify creator")
			}
		}
	}

	return rcv.SendMessage(chatID, message)
}
