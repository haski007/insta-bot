package listener

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/haski007/insta-bot/pkg/emoji"
)

func (rcv *InstaBotService) cmdSubARCEventHandler(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	if err := rcv.storage.SubscribeChatToARC(chatID); err != nil {
		rcv.log.WithError(err).Error("[cmdSubARCEventHandler] subscribe chat to ARC")
		return
	}

	if err := rcv.SendMessage(chatID, "You have successfully subscribed to ARC events"+emoji.Check); err != nil {
		rcv.log.WithError(err).Error("[cmdSubARCEventHandler] send message")
		rcv.NotifyCreator(fmt.Sprintf("[cmdSubARCEventHandler] send message: %s\n", err))
		return
	}
}

func (rcv *InstaBotService) cmdUnsubARCEventHandler(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	if err := rcv.storage.UnsubscribeChatToARC(chatID); err != nil {
		rcv.log.WithError(err).Error("[cmdUnsubARCEventHandler] unsubscribe chat to ARC")
		return
	}

	if err := rcv.SendMessage(chatID, "You have successfully unsubscribed from ARC events"+emoji.Basket); err != nil {
		rcv.log.WithError(err).Error("[cmdUnsubARCEventHandler] send message")
		rcv.NotifyCreator(fmt.Sprintf("[cmdUnsubARCEventHandler] send message: %s\n", err))
		return
	}
}