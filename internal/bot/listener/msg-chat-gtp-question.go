package listener

import (
	"context"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (rcv *InstaBotService) msgChatGPTQuestion(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	messageID := update.Message.MessageID

	question := strings.TrimPrefix(update.Message.Text, "?")
	rsp, err := rcv.gpt.GetShortAnswer(context.Background(), question)
	if err != nil {
		rcv.log.WithError(err).Error("[msgChatGPTQuestion] get short answer")
		rcv.SendError(chatID, ErrInternalServerError)
		rcv.NotifyCreator(fmt.Sprintf("[msgChatGPTQuestion] get short answer: %s\n", err))
		return
	}

	if err := rcv.Reply(chatID, messageID, rsp); err != nil {
		rcv.SendError(chatID, ErrInternalServerError)
		rcv.NotifyCreator(fmt.Sprintf("[msgChatGPTQuestion] send message: %s\n", err))
		return
	}
}
