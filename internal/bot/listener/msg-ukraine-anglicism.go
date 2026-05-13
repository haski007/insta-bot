package listener

import (
	"context"
	"strings"
	"time"
	"unicode/utf8"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (rcv *InstaBotService) msgUkraineAnglicismIfNeeded(update tgbotapi.Update) {
	if rcv.openRouter == nil || update.Message == nil {
		return
	}
	msg := update.Message
	if msg.From != nil && msg.From.IsBot {
		return
	}

	text := strings.TrimSpace(msg.Text)
	if text == "" {
		text = strings.TrimSpace(msg.Caption)
	}
	if utf8.RuneCountInString(text) < 4 {
		return
	}

	subscribed, err := rcv.storage.IsChatSubscribedToUkraineForUkrainians(msg.Chat.ID)
	if err != nil {
		rcv.log.WithError(err).Error("[msgUkraineAnglicismIfNeeded] IsChatSubscribedToUkraineForUkrainians")
		return
	}
	if !subscribed {
		return
	}

	if utf8.RuneCountInString(text) > 3500 {
		r := []rune(text)
		text = string(r[:3500])
	}

	ctx, cancel := context.WithTimeout(rcv.ctx, 50*time.Second)
	defer cancel()

	res, err := rcv.openRouter.AnalyzeAnglicisms(ctx, text)
	if err != nil {
		rcv.log.WithError(err).Warn("[msgUkraineAnglicismIfNeeded] AnalyzeAnglicisms")
		return
	}
	if res == nil || !res.HasAnglicism || res.Rewritten == "" {
		return
	}

	out := res.Rewritten
	if err := rcv.ReplyPlain(msg.Chat.ID, msg.MessageID, out); err != nil {
		rcv.log.WithError(err).Error("[msgUkraineAnglicismIfNeeded] ReplyPlain")
		_ = rcv.NotifyCreator("[msgUkraineAnglicismIfNeeded] ReplyPlain: " + err.Error())
	}
}
