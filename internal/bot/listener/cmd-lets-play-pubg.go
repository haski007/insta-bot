package listener

import (
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (rcv *InstaBotService) cmdLetsPlayPUBGHandler(update tgbotapi.Update) {
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
	chatRegistered, err := rcv.storage.ChatExists(PUBGContext, chatID)
	if err != nil {
		rcv.SendError(chatID, ErrInternalServerError)
		rcv.log.WithError(err).Error("[cmdLetsPlayPUBGHandler] get chat exists")
		return
	}
	if !chatRegistered {
		rcv.SendError(chatID, ErrNoPUBGPlayers)
		return
	}

	members, err := rcv.storage.GetChatMembers(PUBGContext, chatID)
	if err != nil {
		rcv.SendError(chatID, ErrInternalServerError)
		rcv.log.WithError(err).Error("[cmdLetsPlayPUBGHandler] get chat members")
		return
	}

	if len(members) == 0 {
		rcv.SendError(chatID, ErrNoPUBGPlayers)
		return
	}

	var message string
	for _, m := range members {
		message += fmt.Sprintf("@%s ", m)
	}

	var voteCaption = "Галасаваніє!"

	if timeToPlay != "" {
		voteCaption = fmt.Sprintf("%s PUBG в %s?", voteCaption, timeToPlay)
		message += fmt.Sprintf("\nХто буде в PUBG в %s? Галасуєм!", timeToPlay)
	} else {
		message += "\nХто буде в PUBG? Галасуєм!"
	}
	message = strings.ReplaceAll(message, "_", "\\_")

	if err := rcv.SendMessageWithoutMarkdown(chatID, message); err != nil {
		logrus.WithError(err).Println("[cmdLetsPlayPUBGHandler] send message to chat")
	}

	options := []string{
		"(+) Звичайно, я справжній козак!",
		"(-) Нііі, ні я не ту кохав, не ті слова..., в общем лох я!",
	}

	if err := rcv.CreatePoll(chatID, voteCaption, false, options...); err != nil {
		logrus.WithError(err).Println("[cmdLetsPlayPUBGHandler] create poll to chat")
	}
}
