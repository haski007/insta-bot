package listener

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func (rcv *InstaBotService) msgSaveToHistory(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	messageID := update.Message.MessageID
	username := update.Message.From.UserName
	text := update.Message.Text

	if !update.Message.From.IsBot {
		var title string
		if update.Message.Chat.Title == "" {
			title = update.Message.Chat.FirstName + " " + update.Message.Chat.LastName
		} else {
			title = update.Message.Chat.Title
		}

		logrus.WithField("text", text).
			WithField("chatID", chatID).
			WithField("chat_title", title).
			WithField("from", username).
			Info("message was sent")
	}

	var fromBlock string
	if update.Message.ForwardFrom != nil {
		fromBlock = fmt.Sprintf("[forwarded from: @%s by @%s]", update.Message.ForwardFrom.UserName, username)
	} else {
		fromBlock = fmt.Sprintf("[from: @%s]", username)
	}

	if update.Message.From.IsBot || update.Message.IsCommand() {
		return
	}

	if err := rcv.storage.SaveMessage(
		chatID,
		messageID,
		fmt.Sprintf("%s, message: ```%s```", fromBlock, text),
	); err != nil {
		rcv.log.WithError(err).Error("[msgSaveToHistory] save message")
		rcv.NotifyCreator(fmt.Sprintf("[msgSaveToHistory] save message: %s\n", err))
		return
	}
}
