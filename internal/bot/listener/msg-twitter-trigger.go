package listener

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (rcv *InstaBotService) msgTwitterTrigger(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	messageID := update.Message.MessageID
	url := exprFindURL.FindString(update.Message.Text)

	url = strings.ReplaceAll(url, "https://x.com/", "https://vxtwitter.com/")

	if err := rcv.SendMessageWithoutMarkdown(chatID, fmt.Sprintf("forwarder: @%s\n\nurl: %s", update.Message.From.UserName, url)); err != nil {
		rcv.log.WithError(err).Error("[msgTwitterTrigger] send message caption")
		return
	}

	if err := rcv.DeleteMessage(chatID, messageID); err != nil {
		rcv.log.WithError(err).Error("[msgTwitterTrigger] delete message")
		return
	}
}
