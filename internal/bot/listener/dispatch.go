package listener

import (
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/haski007/insta-bot/internal/metrics"
	"github.com/haski007/insta-bot/pkg/emoji"
	"github.com/haski007/insta-bot/pkg/safego"
)

func updateKind(update tgbotapi.Update) string {
	switch {
	case update.EditedMessage != nil:
		return "edited"
	case update.Poll != nil:
		return "poll"
	case update.PollAnswer != nil:
		return "poll_answer"
	case update.MyChatMember != nil:
		return "my_chat_member"
	case update.Message != nil && update.Message.IsCommand():
		return "command"
	case update.Message != nil:
		return "message"
	default:
		return "other"
	}
}

func (rcv *InstaBotService) goCommand(name string, update tgbotapi.Update, fn func(tgbotapi.Update)) {
	metrics.IncRoute("command")
	go func() {
		start := time.Now()
		fn(update)
		metrics.ObserveCommand(name, start)
	}()
}

func (rcv *InstaBotService) goCommandSafe(name string, update tgbotapi.Update, fn func(tgbotapi.Update)) {
	metrics.IncRoute("command")
	start := time.Now()
	safego.New(func() {
		defer metrics.ObserveCommand(name, start)
		fn(update)
	}, func(pErr any) {
		rcv.log.WithError(fmt.Errorf("%s", pErr)).Errorf("[%s] panic", name)
		rcv.NotifyCreator(fmt.Sprintf("%s %s panic: %s", emoji.NoEntry, name, pErr))
	})
}

func (rcv *InstaBotService) goMessage(kind string, update tgbotapi.Update, fn func(tgbotapi.Update)) {
	metrics.IncRoute(kind)
	go func() {
		start := time.Now()
		fn(update)
		metrics.ObserveMessageHandler(kind, start)
	}()
}
