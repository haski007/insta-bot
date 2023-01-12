package listener

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (rcv *InstaBotService) SendMessage(chatID int64, text string) error {
	message := tgbotapi.NewMessage(chatID, text)
	message.ParseMode = tgbotapi.ModeMarkdown

	_, err := rcv.bot.Send(message)
	return err
}

func (rcv *InstaBotService) SendMessageWithoutMarkdown(chatID int64, text string) error {
	message := tgbotapi.NewMessage(chatID, text)

	_, err := rcv.bot.Send(message)
	return err
}

func (rcv *InstaBotService) CreatePoll(
	chatID int64, caption string, isAnon bool,
	options ...string,
) error {
	answer := tgbotapi.NewPoll(chatID, caption, options...)
	answer.IsAnonymous = isAnon
	_, err := rcv.bot.Send(answer)
	return err
}
