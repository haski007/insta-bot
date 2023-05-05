package listener

import (
	"github.com/haski007/insta-bot/pkg/emoji"
	"github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (rcv *InstaBotService) cmdTestHandler(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	//msg := tgbotapi.NewMessage(chatID, "Please pay for your purchase")

	invoice := tgbotapi.NewInvoice(chatID,
		"Purchase",
		"Purchase Description",
		"payload",
		"provider_token",
		"start_parameter",
		"USD",
		[]tgbotapi.LabeledPrice{
			{
				Label:  "Health course",
				Amount: 0,
			},
		},
	)
	invoice.MaxTipAmount = 12

	if err := rcv.SendMessage(chatID, "Pososi! "+emoji.Gear); err != nil {
		logrus.WithError(err).Println("[cmdTestHandler] send message to chat")
	}

	if _, err := rcv.bot.Send(invoice); err != nil {
		logrus.WithError(err).Println("send invoice")
	}
}
