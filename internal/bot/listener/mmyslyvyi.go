package listener

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const mmyslyvyiUsername = "mmyslyvyi"

const ukraineAnglicismMmyslyvyiSuffix = "\n(Автор повідомлення хуєсос)"

const mmyslyvyiPoopReaction = "💩"

type reactionTypeEmoji struct {
	Type  string `json:"type"`
	Emoji string `json:"emoji"`
}

func isMmyslyvyi(username string) bool {
	un := strings.TrimPrefix(strings.ToLower(strings.TrimSpace(username)), "@")
	return un == mmyslyvyiUsername
}

// ReactPoop sets 💩 on a message (setMessageReaction; not in tgbotapi v5.5.1).
func (rcv *InstaBotService) ReactPoop(chatID int64, messageID int) error {
	params := make(tgbotapi.Params)
	if err := params.AddFirstValid("chat_id", chatID); err != nil {
		return err
	}
	params.AddNonZero("message_id", messageID)
	if err := params.AddInterface("reaction", []reactionTypeEmoji{{
		Type:  "emoji",
		Emoji: mmyslyvyiPoopReaction,
	}}); err != nil {
		return err
	}
	_, err := rcv.bot.MakeRequest("setMessageReaction", params)
	return err
}

// isMessageAlive returns true if the message still exists in the chat. It
// performs a no-op setMessageReaction (clear reactions) — if the API returns
// "message not found" the message has been deleted. Anything else is treated
// as "alive" to avoid false-positive escalations.
func (rcv *InstaBotService) isMessageAlive(chatID int64, messageID int) bool {
	params := make(tgbotapi.Params)
	if err := params.AddFirstValid("chat_id", chatID); err != nil {
		return true
	}
	params.AddNonZero("message_id", messageID)
	if err := params.AddInterface("reaction", []reactionTypeEmoji{}); err != nil {
		return true
	}
	if _, err := rcv.bot.MakeRequest("setMessageReaction", params); err != nil {
		errStr := strings.ToLower(err.Error())
		if strings.Contains(errStr, "message to react not found") ||
			strings.Contains(errStr, "message not found") ||
			strings.Contains(errStr, "message_id_invalid") ||
			strings.Contains(errStr, "message to edit not found") {
			return false
		}
		rcv.log.WithError(err).Debug("[isMessageAlive] unexpected error, assume alive")
	}
	return true
}

func (rcv *InstaBotService) reactMmyslyvyiIfNeeded(msg *tgbotapi.Message) {
	if msg == nil || msg.From == nil || msg.From.IsBot {
		return
	}
	if !isMmyslyvyi(msg.From.UserName) {
		return
	}
	if err := rcv.ReactPoop(msg.Chat.ID, msg.MessageID); err != nil {
		rcv.log.WithError(err).Debug("[reactMmyslyvyiIfNeeded] setMessageReaction")
	}
}
