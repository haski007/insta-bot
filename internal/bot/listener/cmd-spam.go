package listener

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	maxSpamNum = 200
)

func (rcv *InstaBotService) cmdSpam(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	if !rcv.IsCreator(update.Message.From.ID) {
		rcv.SendError(chatID, ErrAccessDenied)
	}

	args := strings.Fields(update.Message.CommandArguments())

	if len(args) < 1 {
		rcv.SendError(chatID, ErrNoArguments)
		return
	}

	num, err := strconv.Atoi(args[0])
	if err != nil {
		rcv.SendError(chatID, ErrWrongFormat)
		return
	}

	if num > maxSpamNum {
		num = maxSpamNum
	}

	usersToSpam := args[1:]

	for i := 0; i < num; i++ {
		var message string

		for _, user := range usersToSpam {
			if strings.HasPrefix(user, "@") {
				message += fmt.Sprintf("%s ", user)
			} else {
				message += fmt.Sprintf("@%s ", user)
			}
		}

		if err := rcv.SendMessage(chatID, message); err != nil {
			logrus.WithError(err).Println("[cmdSpam] send message to chat")
		}
		time.Sleep(time.Millisecond * 500)
	}
}
