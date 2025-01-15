package listener

import (
	"fmt"
	"strings"

	"github.com/haski007/insta-bot/internal/bot/publisher"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (rcv *InstaBotService) msgTwitterTrigger(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	messageID := update.Message.MessageID
	url := exprFindURL.FindString(update.Message.Text)

	if strings.Contains(url, "/status/") {
		if strings.Contains(url, publisher.TwitterBaseUrl) {
			url = strings.ReplaceAll(url, publisher.TwitterBaseUrl, publisher.VXTwitterBaseUrl)
		} else if strings.Contains(url, publisher.TwitterOLDBaseUrl) {
			url = strings.ReplaceAll(url, publisher.TwitterOLDBaseUrl, publisher.VXTwitterBaseUrl)
		}
	} else {
		return
	}

	if err := rcv.SendMessageWithoutMarkdown(chatID, fmt.Sprintf("forwarder: @%s\n\nurl: %s", update.Message.From.UserName, url)); err != nil {
		rcv.log.WithError(err).Error("[msgTwitterTrigger] send message caption")
		return
	}

	if err := rcv.DeleteMessage(chatID, messageID); err != nil {
		rcv.log.WithError(err).Error("[msgTwitterTrigger] delete message")
		return
	}
}
