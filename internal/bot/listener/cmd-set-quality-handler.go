package listener

import (
	"strconv"

	"github.com/haski007/insta-bot/pkg/emoji"
	"github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (rcv *InstaBotService) cmdSetQualityHandler(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	if update.Message.From.ID != rcv.creatorID {
		rcv.SendError(chatID, "Only creator can do this")
		return
	}

	args := update.Message.CommandArguments()
	if len(args) == 0 {
		rcv.SendError(chatID, "need argument")
		return
	}

	quality, err := strconv.Atoi(args)
	if err != nil {
		rcv.SendError(chatID, "cannot parse your argument to an integer value")
		return
	}

	if quality < 144 {
		rcv.SendError(chatID, "quality cannot be less than 144")
	}

	rcv.youtubeApi.SetMaxQuality(quality)
	rcv.log.WithField("new_max_quality", rcv.youtubeApi.GetMaxQuality()).Println("Set new max quality value")

	if err := rcv.SendMessage(chatID, "New quality value has been set"+emoji.Check); err != nil {
		logrus.WithError(err).Println("[cmdTestHandler] send message to chat")
	}
}
