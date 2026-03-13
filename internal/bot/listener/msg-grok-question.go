package listener

import (
	"context"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (rcv *InstaBotService) msgGrokQuestion(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	messageID := update.Message.MessageID

	if rcv.grokSrv == nil {
		rcv.SendError(chatID, ErrInternalServerError)
		rcv.NotifyCreator("[msgGrokQuestion] Grok service is not configured (missing GROK_API_KEY)")
		return
	}

	question := strings.TrimPrefix(update.Message.Text, "g?")
	rsp, err := rcv.grokSrv.GetShortAnswer(context.Background(), question)
	if err != nil {
		rcv.log.WithError(err).Error("[msgGrokQuestion] get short answer")
		rcv.SendError(chatID, ErrInternalServerError)
		rcv.NotifyCreator(fmt.Sprintf("[msgGrokQuestion] get short answer: %s\n", err))
		return
	}

	if err := rcv.Reply(chatID, messageID, rsp); err != nil {
		rcv.SendError(chatID, ErrInternalServerError)
		rcv.NotifyCreator(fmt.Sprintf("[msgGrokQuestion] send message: %s\n", err))
		return
	}
}
