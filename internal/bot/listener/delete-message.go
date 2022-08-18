package listener

import (
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (rcv *InstaBotService) DeleteMessage(chatID int64, messageID int) error {
	req := tgbotapi.NewDeleteMessage(chatID, messageID)

	_, err := rcv.bot.Send(req)
	if err != nil {
		if _, ok := err.(*json.UnmarshalTypeError); !ok {
			return err
		}
	}

	return nil
}
