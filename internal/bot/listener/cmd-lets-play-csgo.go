package listener

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/haski007/insta-bot/internal/clients/google"
	"github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (rcv *InstaBotService) cmdLetsPlayHandler(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	args := strings.Fields(update.Message.CommandArguments())

	var (
		timeToPlay    string
		suggestedTime time.Time
		err           error
	)
	if len(args) != 0 {
		suggestedTime, err = time.Parse("15:04", args[0])
		if err == nil {
			timeToPlay = args[0]
		}
		suggestedTime = time.Date(time.Now().Year(),
			time.Now().Month(),
			time.Now().Day(),
			suggestedTime.Hour(),
			suggestedTime.Minute(),
			suggestedTime.Second(),
			suggestedTime.Nanosecond(),
			time.Now().Location())
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

		rsp, err := rcv.calendar.CreateMeet(context.Background(), &google.CreateMeetReq{
			Summary:     "Very important meeting",
			Location:    "Forsage dance club",
			Description: "Here we gonna make some magic and win all the enemies!",
			CalendarID:  "primary",
			Guests:      nil,
			StartTime:   suggestedTime,
			EndTime:     suggestedTime.Add(time.Hour),
		})
		if err != nil {
			rcv.NotifyCreator(fmt.Sprintf("can't create google meet err: %s", err))
			rcv.log.WithError(err).Println("create google meet")
		}

		time.AfterFunc(suggestedTime.Sub(time.Now()), func() {
			rcv.SendMessage(chatID, "Here we go guys! "+rsp.MeetLink)
		})
	} else {
		message += "\nХто буде в коунтер стріке? Галасуєм!"
	}

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
