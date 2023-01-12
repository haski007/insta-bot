package listener

import (
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (rcv *InstaBotService) cmdLetsPlayHandler(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	args := strings.Fields(update.Message.CommandArguments())

	var timeToPlay string
	if len(args) != 0 {
		_, err := time.Parse("15:04", args[0])
		if err == nil {
			timeToPlay = args[0]
		}
	}

	// ---> Check if chat is registered not to spam in usual chats
	chatRegistered, err := rcv.storage.ChatExists(CSGOContext, chatID)
	if err != nil {
		rcv.SendError(chatID, ErrInternalServerError)
		rcv.log.WithError(err).Error("[cmdLetsPlayHandler] get chat exists")
		return
	}
	if !chatRegistered {
		rcv.SendError(chatID, ErrNoCSGOPlayers)
		return
	}

	members, err := rcv.storage.GetChatMembers(CSGOContext, chatID)
	if err != nil {
		rcv.SendError(chatID, ErrInternalServerError)
		rcv.log.WithError(err).Error("[cmdLetsPlayHandler] get chat members")
		return
	}

	if len(members) == 0 {
		rcv.SendError(chatID, ErrNoCSGOPlayers)
		return
	}

	var message string
	for _, m := range members {
		message += fmt.Sprintf("@%s ", m)
	}

	var voteCaption = "Галасаваніє!"

	if timeToPlay != "" {
		voteCaption = fmt.Sprintf("%s CS GO в %s?", voteCaption, timeToPlay)
		message += fmt.Sprintf("\nХто буде в коунтер стріке в %s? Галасуєм!", timeToPlay)
	} else {
		message += "\nХто буде в коунтер стріке? Галасуєм!"
	}
	strings.ReplaceAll(message, "_", "\\_")

	if err := rcv.SendMessageWithoutMarkdown(chatID, message); err != nil {
		logrus.WithError(err).Println("[cmdLetsPlayHandler] send message to chat")
	}

	options := []string{
		"(+) Звичайно, я справжній козак!",
		"(-) Нііі, ні я не ту кохав, не ті слова..., в общем лох я!",
	}

	if err := rcv.CreatePoll(chatID, voteCaption, false, options...); err != nil {
		logrus.WithError(err).Println("[cmdLetsPlayHandler] create poll to chat")
	}
}
