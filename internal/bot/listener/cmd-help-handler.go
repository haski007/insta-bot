package listener

import (
	"github.com/haski007/insta-bot/pkg/emoji"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (rcv *InstaBotService) cmdStartHandler(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	if err := rcv.sendStartInfo(chatID); err != nil {
		rcv.log.WithError(err).Println("[cmdStartHandler] send start info")
	}
}

func (rcv *InstaBotService) sendStartInfo(chatID int64) error {
	message := `Hello, I am a bot that can help you make your chat meme life a lot easier and funnier.
For the best user experience give me please administrator rights ` + emoji.FaceWinking + "\n" +
		"you can use such commands for creating csgo polls:\n" +
		"/reg\\_csgo\\_players {username\\_1} {username\\_2} ... - to register users that will be tagged before poll to play CS GO\n" +
		"/purge\\_csgo\\_players - to delete all csgo players at this chat\n" +
		"/lets\\_play - to offer to play CS GO (creates poll to vote)\n" +
		"/sum {count} - to summarize last {count > 0} messages\n" +
		"/purge\\_history - to delete all messages from this chat history\n"
	return rcv.SendMessage(chatID, message)
}
