package listener

import (
	"fmt"
	"math/rand"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// IG-bait params: post a brief Instagram reel URL (hidden behind a spoiler) so
// any chat-wide IG-aware filter (e.g., Dyoma bot's instagram_preceded) marks
// this chat as "IG just appeared" and lets our subsequent video reply through.
// The bait is removed shortly after the real reply lands.
const (
	anglicismBaitWarmup            = 1100 * time.Millisecond
	anglicismBaitCleanupAfter      = 5 * time.Second
	anglicismBaitFastCleanup       = 200 * time.Millisecond
	anglicismBaitShortcodeAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_-"
	anglicismBaitShortcodeLen      = 11
)

// buildAnglicismBaitMessage returns a Telegram HTML message with a spoiler-wrapped
// Instagram reel URL — the URL is in raw text (matches enemy regex) but visually
// it's a tiny gray block in chat.
func buildAnglicismBaitMessage() string {
	b := make([]byte, anglicismBaitShortcodeLen)
	for i := range b {
		b[i] = anglicismBaitShortcodeAlphabet[rand.Intn(len(anglicismBaitShortcodeAlphabet))]
	}
	return fmt.Sprintf(`<tg-spoiler>https://www.instagram.com/reel/%s/</tg-spoiler>`, string(b))
}

// sendInstagramBait posts the bait message and returns its ID for later deletion.
func (rcv *InstaBotService) sendInstagramBait(chatID int64) (int, error) {
	msg := tgbotapi.NewMessage(chatID, buildAnglicismBaitMessage())
	msg.ParseMode = tgbotapi.ModeHTML
	msg.DisableWebPagePreview = true
	msg.DisableNotification = true
	sent, err := rcv.bot.Send(msg)
	if err != nil {
		return 0, err
	}
	return sent.MessageID, nil
}

// deleteMessageAfter sleeps for delay then deletes the message (best-effort).
func (rcv *InstaBotService) deleteMessageAfter(chatID int64, messageID int, delay time.Duration) {
	time.Sleep(delay)
	if err := rcv.DeleteMessage(chatID, messageID); err != nil {
		rcv.log.WithError(err).Debug("[deleteMessageAfter] DeleteMessage")
	}
}
