package listener

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/haski007/insta-bot/pkg/emoji"
)

func (rcv *InstaBotService) cmdUkraineForUkrainiansSub(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	if err := rcv.storage.SubscribeChatToUkraineForUkrainians(chatID); err != nil {
		rcv.log.WithError(err).Error("[cmdUkraineForUkrainiansSub] subscribe")
		_ = rcv.SendMessageWithoutMarkdown(chatID, "could not subscribe to ukraine for ukrainians"+emoji.NoEntry)
		return
	}

	// Plain text: command contains underscores which break legacy Markdown.
	if err := rcv.SendMessageWithoutMarkdown(chatID, "Увімкнено моніторинг англіцизмів у цьому чаті. Вимкнути: /unsub_ukraine_for_ukrainians"+emoji.Check); err != nil {
		rcv.log.WithError(err).Error("[cmdUkraineForUkrainiansSub] send message")
		_ = rcv.NotifyCreator(fmt.Sprintf("[cmdUkraineForUkrainiansSub] send: %s\n", err))
	}
}

func (rcv *InstaBotService) cmdUkraineForUkrainiansUnsub(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	if err := rcv.storage.UnsubscribeChatFromUkraineForUkrainians(chatID); err != nil {
		rcv.log.WithError(err).Error("[cmdUkraineForUkrainiansUnsub] unsubscribe")
		_ = rcv.SendMessageWithoutMarkdown(chatID, "could not unsubscribe from ukraine for ukrainians"+emoji.NoEntry)
		return
	}

	if err := rcv.SendMessageWithoutMarkdown(chatID, "Моніторинг англіцизмів вимкнено"+emoji.Basket); err != nil {
		rcv.log.WithError(err).Error("[cmdUkraineForUkrainiansUnsub] send message")
		_ = rcv.NotifyCreator(fmt.Sprintf("[cmdUkraineForUkrainiansUnsub] send: %s\n", err))
	}
}
