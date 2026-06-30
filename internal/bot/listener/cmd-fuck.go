package listener

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/haski007/insta-bot/pkg/emoji"
)

func (rcv *InstaBotService) isInFuckList(username string) bool {
	user, ok := normalizeUkraineIgnoreUsername(username)
	if !ok {
		return false
	}
	inList, err := rcv.storage.FuckListContains(user)
	if err != nil {
		rcv.log.WithError(err).Debug("[isInFuckList] FuckListContains")
		return false
	}
	return inList
}

func (rcv *InstaBotService) cmdFuck(update tgbotapi.Update) {
	if !rcv.IsCreator(update.Message.From.ID) {
		rcv.SendMessage(update.Message.Chat.ID, "You are not allowed to use this command"+emoji.NoEntry)
		return
	}

	args := rcv.parseCommandArgs(update.Message.Text)
	if len(args) < 1 {
		rcv.SendMessage(update.Message.Chat.ID, "Usage: /fuck @username")
		return
	}

	user, ok := normalizeUkraineIgnoreUsername(args[0])
	if !ok {
		rcv.SendMessage(update.Message.Chat.ID, "Invalid username. Example: /fuck @username"+emoji.NoEntry)
		return
	}

	if err := rcv.storage.FuckListAdd(user); err != nil {
		rcv.log.WithError(err).Error("[cmdFuck] FuckListAdd")
		rcv.SendMessage(update.Message.Chat.ID, "Failed to update fuck list"+emoji.NoEntry)
		return
	}

	rcv.SendMessage(update.Message.Chat.ID, fmt.Sprintf("Added to fuck list: @%s 💩", user)+emoji.Check)
}

func (rcv *InstaBotService) cmdUnfuck(update tgbotapi.Update) {
	if !rcv.IsCreator(update.Message.From.ID) {
		rcv.SendMessage(update.Message.Chat.ID, "You are not allowed to use this command"+emoji.NoEntry)
		return
	}

	args := rcv.parseCommandArgs(update.Message.Text)
	if len(args) < 1 {
		rcv.SendMessage(update.Message.Chat.ID, "Usage: /unfuck @username")
		return
	}

	user, ok := normalizeUkraineIgnoreUsername(args[0])
	if !ok {
		rcv.SendMessage(update.Message.Chat.ID, "Invalid username. Example: /unfuck @username"+emoji.NoEntry)
		return
	}

	inList, err := rcv.storage.FuckListContains(user)
	if err != nil {
		rcv.log.WithError(err).Error("[cmdUnfuck] FuckListContains")
		rcv.SendMessage(update.Message.Chat.ID, "Failed to check fuck list"+emoji.NoEntry)
		return
	}
	if !inList {
		rcv.SendMessage(update.Message.Chat.ID, "User not found in fuck list"+emoji.NoEntry)
		return
	}

	if err := rcv.storage.FuckListRemove(user); err != nil {
		rcv.log.WithError(err).Error("[cmdUnfuck] FuckListRemove")
		rcv.SendMessage(update.Message.Chat.ID, "Failed to update fuck list"+emoji.NoEntry)
		return
	}

	rcv.SendMessage(update.Message.Chat.ID, fmt.Sprintf("Removed from fuck list: @%s", user)+emoji.Check)
}
