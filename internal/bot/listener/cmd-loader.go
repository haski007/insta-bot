package listener

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/haski007/insta-bot/pkg/emoji"
)

func (rcv *InstaBotService) cmdDisableLoaderHandler(update tgbotapi.Update) {
	chatId := update.Message.Chat.ID
	err := rcv.storage.DisableLoaderForChat(chatId)
	if err != nil {
		rcv.log.WithError(err).Error("cmdDisableLoaderHandler")
		return
	}
	rcv.SendMessage(chatId, "Loader disabled " + emoji.NoEntry)
}

func (rcv *InstaBotService) cmdEnableLoaderHandler(update tgbotapi.Update) {
	chatId := update.Message.Chat.ID
	err := rcv.storage.EnableLoaderForChat(chatId)
	if err != nil {
		rcv.log.WithError(err).Error("cmdEnableLoaderHandler")
		return
	}
	rcv.SendMessage(chatId, "Loader enabled " + emoji.Check)
}