package listener

import (
	"context"
	"html"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// LLM returns markers like {{українське}}(було: англіцизм); we turn them into Telegram HTML.
var ukraineAnglicismMarker = regexp.MustCompile(`\{\{([^}]+)\}\}\s*\(\s*було\s*:\s*([^)]+)\)`)

func formatAnglicismRewrittenAsTelegramHTML(rewritten string) string {
	var b strings.Builder
	last := 0
	for _, loc := range ukraineAnglicismMarker.FindAllStringSubmatchIndex(rewritten, -1) {
		before := rewritten[last:loc[0]]
		b.WriteString(html.EscapeString(before))
		newW := rewritten[loc[2]:loc[3]]
		oldW := rewritten[loc[4]:loc[5]]
		b.WriteString("<s>")
		b.WriteString(html.EscapeString(strings.TrimSpace(oldW)))
		b.WriteString("</s> → <b>")
		b.WriteString(html.EscapeString(strings.TrimSpace(newW)))
		b.WriteString("</b>")
		last = loc[1]
	}
	b.WriteString(html.EscapeString(rewritten[last:]))
	return b.String()
}

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

	body := formatAnglicismRewrittenAsTelegramHTML(res.Rewritten)
	out := body
	if utf8.RuneCountInString(out) > 4000 {
		r := []rune(out)
		out = string(r[:3997]) + "…"
	}

	if err := rcv.ReplyHTML(msg.Chat.ID, msg.MessageID, out); err != nil {
		rcv.log.WithError(err).Error("[msgUkraineAnglicismIfNeeded] ReplyHTML")
		if err2 := rcv.ReplyPlain(msg.Chat.ID, msg.MessageID, res.Rewritten); err2 != nil {
			rcv.log.WithError(err2).Error("[msgUkraineAnglicismIfNeeded] ReplyPlain fallback")
			_ = rcv.NotifyCreator("[msgUkraineAnglicismIfNeeded] ReplyHTML: " + err.Error())
		}
	}
}
