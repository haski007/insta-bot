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

// Delay between sending an escalation step and checking if it was deleted by
// the enemy bot.
const anglicismDeleteCheckDelay = 1100 * time.Millisecond

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

	rcv.deliverAnglicismReply(msg.Chat.ID, msg.MessageID, out)
}

// deliverAnglicismReply chooses delivery based on env mode:
//   - "escalate" (default): plain → if deleted, photo → if deleted, video w/ IG bait
//   - "plain"  : plain text only, no escalation
//   - "photo"  : photo card only
//   - "video"  : video card with IG bait only
func (rcv *InstaBotService) deliverAnglicismReply(chatID int64, messageID int, text string) {
	mode := strings.ToLower(strings.TrimSpace(rcv.ukraineAnglicismDelivery))
	if mode == "" {
		mode = "escalate"
	}

	switch mode {
	case "plain":
		rcv.sendAnglicismPlainBestEffort(chatID, text)
	case "photo":
		if !rcv.sendAnglicismPhotoCardBestEffort(chatID, messageID, text) {
			rcv.sendAnglicismPlainBestEffort(chatID, text)
		}
	case "video":
		if !rcv.tryDeliverAnglicismVideoWithBait(chatID, messageID, text) {
			rcv.sendAnglicismPlainBestEffort(chatID, text)
		}
	default:
		rcv.deliverAnglicismEscalating(chatID, messageID, text)
	}
}

// deliverAnglicismEscalating sends plain text; if the enemy deletes it within
// ~1s, escalates to photo, then to video with IG bait. Final fallback — a
// repeat plain send (best-effort) so the chat is not left silent.
func (rcv *InstaBotService) deliverAnglicismEscalating(chatID int64, messageID int, text string) {
	if id, err := rcv.sendAnglicismPlain(chatID, text); err != nil {
		rcv.log.WithError(err).Warn("[deliverAnglicismEscalating] sendAnglicismPlain")
	} else {
		time.Sleep(anglicismDeleteCheckDelay)
		if rcv.isMessageAlive(chatID, id) {
			return
		}
		rcv.log.Info("[deliverAnglicismEscalating] plain deleted, escalating to photo")
	}

	if id, err := rcv.sendAnglicismPhotoCard(chatID, messageID, text); err != nil {
		rcv.log.WithError(err).Warn("[deliverAnglicismEscalating] sendAnglicismPhotoCard")
	} else {
		time.Sleep(anglicismDeleteCheckDelay)
		if rcv.isMessageAlive(chatID, id) {
			return
		}
		rcv.log.Info("[deliverAnglicismEscalating] photo deleted, escalating to video")
	}

	if rcv.tryDeliverAnglicismVideoWithBait(chatID, messageID, text) {
		return
	}

	rcv.sendAnglicismPlainBestEffort(chatID, text)
}

// sendAnglicismPlain sends a plain (non-reply) text message and returns its ID.
func (rcv *InstaBotService) sendAnglicismPlain(chatID int64, text string) (int, error) {
	msg := tgbotapi.NewMessage(chatID, text)
	sent, err := rcv.bot.Send(msg)
	if err != nil {
		return 0, err
	}
	return sent.MessageID, nil
}

func (rcv *InstaBotService) sendAnglicismPlainBestEffort(chatID int64, text string) {
	if _, err := rcv.sendAnglicismPlain(chatID, text); err != nil {
		rcv.log.WithError(err).Error("[sendAnglicismPlainBestEffort] send")
		_ = rcv.NotifyCreator("[sendAnglicismPlainBestEffort] send: " + err.Error())
	}
}

// sendAnglicismPhotoCard renders the text as a card image and sends it as a
// reply to the offending message; returns the sent message ID.
func (rcv *InstaBotService) sendAnglicismPhotoCard(chatID int64, replyToID int, text string) (int, error) {
	cardPNG, err := renderAnglicismCard(text)
	if err != nil {
		return 0, err
	}
	photo := tgbotapi.NewPhoto(chatID, tgbotapi.FileBytes{
		Name:  "card.png",
		Bytes: cardPNG,
	})
	photo.ReplyToMessageID = replyToID
	sent, err := rcv.bot.Send(photo)
	if err != nil {
		return 0, err
	}
	return sent.MessageID, nil
}

func (rcv *InstaBotService) sendAnglicismPhotoCardBestEffort(chatID int64, replyToID int, text string) bool {
	if _, err := rcv.sendAnglicismPhotoCard(chatID, replyToID, text); err != nil {
		rcv.log.WithError(err).Warn("[sendAnglicismPhotoCardBestEffort] send")
		return false
	}
	return true
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
