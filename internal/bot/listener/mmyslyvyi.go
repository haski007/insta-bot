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
