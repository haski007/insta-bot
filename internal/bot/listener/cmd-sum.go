package listener

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (rcv *InstaBotService) cmdSum(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	args := rcv.parseCommandArgs(update.Message.Text)
	if len(args) == 0 {
		rcv.SendError(chatID, ErrNoArguments)
		return
	}

	countOfLastMessages, err := strconv.Atoi(args[0])
	if err != nil || countOfLastMessages < 1 {
		rcv.SendError(chatID, ErrWrongFormat)
		return
	}

	var question string
	if len(args) > 1 {
		question = strings.Join(args[1:], " ")
	}

	messages, err := rcv.storage.GetMessages(chatID, countOfLastMessages)
	if err != nil {
		rcv.SendError(chatID, ErrInternalServerError)
		rcv.log.WithError(err).Error("[cmdSum] get messages")
		return
	}

	summarized, err := rcv.gpt.SummarizeMessages(context.Background(), messages, question)
	if err != nil {
		rcv.SendError(chatID, ErrInternalServerError)
		rcv.log.WithError(err).Error("[cmdSum] summarize messages")
		rcv.NotifyCreator(fmt.Sprintf("[cmdSum] summarize messages err: %s", err))
		return
	}

	if _, err := rcv.bot.Send(tgbotapi.NewMessage(chatID, summarized)); err != nil {
		rcv.SendError(chatID, ErrInternalServerError)
		rcv.log.WithError(err).Error("[cmdSum] send summarized")
		return
	}
}
