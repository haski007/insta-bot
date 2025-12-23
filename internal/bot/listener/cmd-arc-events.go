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

func (rcv *InstaBotService) cmdListArcEventsHandler(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	// Get all events
	events, err := rcv.arcRaidersClient.GetARCEvents(nil)
	if err != nil {
		rcv.log.WithError(err).Error("[cmdListArcEventsHandler] get ARC events")
		return
	}

	// Filter events: Matriarch, Harvester, or Night Raid on Stella Montis
	filteredEvents := rcv.filterARCEvents(events)

	// If no matching events, return early
	if len(filteredEvents) == 0 {
		rcv.log.Error("[checkAndNotifyARCEvents] no matching events (Matriarch/Harvester/Night Raid on Stella Montis)")
		rcv.SendMessage(chatID, "No matching events (Matriarch/Harvester/Night Raid on Stella Montis)")
		return
	}

	// Get all subscribed chats
	chatIDs, err := rcv.storage.GetAllARCSubscribedChats()
	if err != nil {
		rcv.log.WithError(err).Error("[cmdListArcEventsHandler] get all ARC subscribed chats")
		rcv.SendMessage(chatID, "Error getting all ARC subscribed chats")
		return
	}

	if len(chatIDs) == 0 {
		rcv.log.Info("[checkAndNotifyARCEvents] no subscribed chats")
		rcv.SendMessage(chatID, "No subscribed chats")
		return
	}

	message, err := rcv.buildARCEventMessage(filteredEvents)
	if err != nil {
		rcv.log.WithError(err).Error("[cmdListArcEventsHandler] build ARC event message")
		rcv.SendMessage(chatID, "Error building ARC event message")
		return
	}

	if err := rcv.SendMessage(chatID, message); err != nil {
		rcv.log.WithError(err).Error("[cmdListArcEventsHandler] send message")
		rcv.NotifyCreator(fmt.Sprintf("[cmdListArcEventsHandler] send message: %s\n", err))
		return
	}
}