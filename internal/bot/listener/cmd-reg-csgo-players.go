package listener

import (
	"fmt"
	"strings"

	"github.com/haski007/insta-bot/pkg/emoji"
	"github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (rcv *InstaBotService) cmdRegCSGOPlayersHandler(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	members := strings.Fields(update.Message.CommandArguments())
	if len(members) == 0 {
		rcv.SendError(chatID, ErrNoArguments)
		return
	}

	if err := rcv.storage.AddChatWithMembers(CSGOContext, chatID, members); err != nil {
		rcv.SendError(chatID, ErrInternalServerError)
		rcv.log.WithError(err).Error("[cmdRegCSGOPlayersHandler] add chat with members")
		return
	}

	message := fmt.Sprintf("Added listener of CSGO triggers for *%d* users ", len(members))
	if err := rcv.SendMessage(chatID, message+emoji.Check); err != nil {
		logrus.WithError(err).Println("[cmdRegCSGOPlayersHandler] send message to chat")
	}
}
