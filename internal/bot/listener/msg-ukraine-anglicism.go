package listener

import (
	"context"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// LLM returns markers like {{українське}}(було: англіцизм); we keep only the Ukrainian replacement in the reply.
var ukraineAnglicismMarker = regexp.MustCompile(`\{\{([^}]+)\}\}\s*\(\s*було\s*:\s*([^)]+)\)`)

func formatAnglicismRewrittenPlain(rewritten string) string {
	var b strings.Builder
	last := 0
	for _, loc := range ukraineAnglicismMarker.FindAllStringSubmatchIndex(rewritten, -1) {
		before := rewritten[last:loc[0]]
		b.WriteString(before)
		newW := rewritten[loc[2]:loc[3]]
		b.WriteString(strings.TrimSpace(newW))
		last = loc[1]
	}
	b.WriteString(rewritten[last:])
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

	if rcv.ukraineAnglicismMaxInputRunes > 0 && utf8.RuneCountInString(text) > rcv.ukraineAnglicismMaxInputRunes {
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

	if msg.From != nil {
		un := strings.TrimPrefix(strings.ToLower(strings.TrimSpace(msg.From.UserName)), "@")
		if un != "" {
			ignored, err := rcv.storage.UkraineAnglicismIgnoreContains(msg.Chat.ID, un)
			if err != nil {
				rcv.log.WithError(err).Error("[msgUkraineAnglicismIfNeeded] UkraineAnglicismIgnoreContains")
			} else if ignored {
				return
			}
		}
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

	out := formatAnglicismRewrittenPlain(res.Rewritten)
	if msg.From != nil && isMmyslyvyi(msg.From.UserName) {
		out += ukraineAnglicismMmyslyvyiSuffix
	}

	cardPNG, cardErr := renderAnglicismCard(out)
	if cardErr != nil {
		rcv.log.WithError(cardErr).Warn("[msgUkraineAnglicismIfNeeded] renderAnglicismCard, fallback to ReplyPlain")
	} else if err := rcv.ReplyPhoto(msg.Chat.ID, msg.MessageID, cardPNG, ""); err != nil {
		rcv.log.WithError(err).Warn("[msgUkraineAnglicismIfNeeded] ReplyPhoto, fallback to ReplyPlain")
	} else {
		return
	}

	if err := rcv.ReplyPlain(msg.Chat.ID, msg.MessageID, out); err != nil {
		rcv.log.WithError(err).Error("[msgUkraineAnglicismIfNeeded] ReplyPlain")
		_ = rcv.NotifyCreator("[msgUkraineAnglicismIfNeeded] ReplyPlain: " + err.Error())
	}
}
