package listener

import (
	"fmt"
	"strings"

	"github.com/haski007/insta-bot/pkg/emoji"
	"github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (rcv *InstaBotService) cmdSetEmailHandler(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	args := strings.Fields(update.Message.CommandArguments())

	if len(args) < 2 {
		rcv.SendError(chatID, ErrNoArguments)
		return
	}

	username, email := args[0], args[1]

	if err := rcv.storage.AddUser(username, email); err != nil {
		rcv.SendError(chatID, ErrInternalServerError)
		rcv.NotifyCreator(fmt.Sprintf("[cmdSetEmailHandler] add user err: %s", err.Error()))
		return
	}

	rcv.log.WithFields(map[string]interface{}{
		"username": username,
		"email":    email,
	}).Debugln("added user")

	message := fmt.Sprintf("Successfully added email *%s* for user: *%s*",
		strings.ReplaceAll(email, "_", "\\_"),
		strings.ReplaceAll(username, "_", "\\_"))
	if err := rcv.SendMessage(chatID, message+emoji.Check); err != nil {
		logrus.WithError(err).Println("[cmdPurgePUBGPlayersHandler] send message to chat")
	}
}
