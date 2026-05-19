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

	rcv.deliverAnglicismReply(msg.Chat.ID, msg.MessageID, out)
}

// deliverAnglicismReply picks the delivery channel based on env (video|photo|plain)
// and falls back through the chain on failure.
func (rcv *InstaBotService) deliverAnglicismReply(chatID int64, messageID int, text string) {
	mode := strings.ToLower(strings.TrimSpace(rcv.ukraineAnglicismDelivery))
	if mode == "" {
		mode = "video"
	}

	if mode == "video" {
		if rcv.tryDeliverAnglicismVideoWithBait(chatID, messageID, text) {
			return
		}
	}

	if mode == "video" || mode == "photo" {
		if cardPNG, err := renderAnglicismCard(text); err != nil {
			rcv.log.WithError(err).Warn("[deliverAnglicismReply] renderAnglicismCard, fallback to plain")
		} else if err := rcv.ReplyPhoto(chatID, messageID, cardPNG, ""); err != nil {
			rcv.log.WithError(err).Warn("[deliverAnglicismReply] ReplyPhoto, fallback to plain")
		} else {
			return
		}
	}

	if err := rcv.ReplyPlain(chatID, messageID, text); err != nil {
		rcv.log.WithError(err).Error("[deliverAnglicismReply] ReplyPlain")
		_ = rcv.NotifyCreator("[deliverAnglicismReply] ReplyPlain: " + err.Error())
	}
}

// tryDeliverAnglicismVideoWithBait posts an "instagram.com" bait, waits for the
// enemy bot to cache it, sends the video card, then schedules bait cleanup.
// Returns true if the video reply was sent successfully.
func (rcv *InstaBotService) tryDeliverAnglicismVideoWithBait(chatID int64, messageID int, text string) bool {
	var baitID int
	if id, err := rcv.sendInstagramBait(chatID); err != nil {
		rcv.log.WithError(err).Warn("[tryDeliverAnglicismVideoWithBait] sendInstagramBait")
	} else {
		baitID = id
		time.Sleep(anglicismBaitWarmup)
	}

	sentOK := false
	if cardMP4, err := renderAnglicismCardVideo(text); err != nil {
		rcv.log.WithError(err).Warn("[tryDeliverAnglicismVideoWithBait] renderAnglicismCardVideo")
	} else if err := rcv.ReplyVideoBytes(chatID, messageID, cardMP4, ""); err != nil {
		rcv.log.WithError(err).Warn("[tryDeliverAnglicismVideoWithBait] ReplyVideoBytes")
	} else {
		sentOK = true
	}

	if baitID != 0 {
		cleanupDelay := anglicismBaitCleanupAfter
		if !sentOK {
			cleanupDelay = anglicismBaitFastCleanup
		}
		go rcv.deleteMessageAfter(chatID, baitID, cleanupDelay)
	}

	return sentOK
}
