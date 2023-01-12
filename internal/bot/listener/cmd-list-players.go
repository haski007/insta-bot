package listener

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (rcv *InstaBotService) cmdListPlayersHandler(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	ctxMembers, err := rcv.storage.GetAllCTXs(chatID)
	if err != nil {
		rcv.SendError(chatID, ErrInternalServerError)
		rcv.log.WithError(err).Error("[cmdListPlayersHandler] GetAllCTXs")
		return
	}

	if len(ctxMembers) == 0 {
		rcv.SendError(chatID, ErrNoPlayers)
		return
	}

	var message string

	for key, members := range ctxMembers {
		game := strings.Split(key, "/")[len(strings.Split(key, "/"))-1]
		message += fmt.Sprintf(
			"_%s_:\n*%s*\n",
			game,
			strings.ReplaceAll(strings.Join(members, "*, *"),
				"_", "\\_"))
	}

	if err := rcv.SendMessage(chatID, message); err != nil {
		logrus.WithError(err).Println("[cmdListPlayersHandler] send message to chat")
	}
}
