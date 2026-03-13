package listener

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/haski007/insta-bot/pkg/emoji"
)

var whoToFuck []string

func (rcv *InstaBotService) cmdFuck(update tgbotapi.Update) {
	if !rcv.IsCreator(update.Message.From.ID) {
		rcv.SendMessage(update.Message.Chat.ID, "You are not allowed to use this command" + emoji.NoEntry)
		return
	}

	args := rcv.parseCommandArgs(update.Message.Text)
	if len(args) < 1 {
		rcv.SendMessage(update.Message.Chat.ID, "Usage: /fuck <user_id>")
		return
	}


	whoToFuck = append(whoToFuck, strings.TrimPrefix(args[0], "@"))
	rcv.SendMessage(update.Message.Chat.ID, "Added to fuck list: "+args[0] + emoji.Check)
}

func (rcv *InstaBotService) cmdUnfuck(update tgbotapi.Update) {
	if !rcv.IsCreator(update.Message.From.ID) {
		rcv.SendMessage(update.Message.Chat.ID, "You are not allowed to use this command" + emoji.NoEntry)
		return
	}

	args := rcv.parseCommandArgs(update.Message.Text)
	if len(args) < 1 {
		rcv.SendMessage(update.Message.Chat.ID, "Usage: /unfuck <user_id>")
		return
	}

	user := strings.TrimPrefix(args[0], "@")

	for i, u := range whoToFuck {
		if u == user {
			whoToFuck = append(whoToFuck[:i], whoToFuck[i+1:]...)
			rcv.SendMessage(update.Message.Chat.ID, "Removed from fuck list: "+args[0] + emoji.Check)
			return
		}
	}
	rcv.SendMessage(update.Message.Chat.ID, "User not found in fuck list" + emoji.NoEntry)
}