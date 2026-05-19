package listener

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (rcv *InstaBotService) SendMessage(chatID int64, text string) error {
	message := tgbotapi.NewMessage(chatID, text)
	message.ParseMode = tgbotapi.ModeMarkdown

	_, err := rcv.bot.Send(message)
	return err
}

func (rcv *InstaBotService) Reply(chatID int64, messageID int, text string) error {
	message := tgbotapi.NewMessage(chatID, text)
	message.ParseMode = tgbotapi.ModeMarkdown
	message.ReplyToMessageID = messageID

	_, err := rcv.bot.Send(message)
	return err
}

// ReplyPlain replies without Markdown to avoid breaking on user punctuation.
func (rcv *InstaBotService) ReplyPlain(chatID int64, messageID int, text string) error {
	message := tgbotapi.NewMessage(chatID, text)
	message.ReplyToMessageID = messageID
	_, err := rcv.bot.Send(message)
	return err
}

// ReplyPhoto replies with an image (raw PNG/JPEG bytes) and optional caption.
func (rcv *InstaBotService) ReplyPhoto(chatID int64, messageID int, imgBytes []byte, caption string) error {
	photo := tgbotapi.NewPhoto(chatID, tgbotapi.FileBytes{
		Name:  "card.png",
		Bytes: imgBytes,
	})
	photo.Caption = caption
	photo.ReplyToMessageID = messageID
	_, err := rcv.bot.Send(photo)
	return err
}

// ReplyHTML sends a reply with Telegram HTML parse mode (escape user text before wrapping in tags).
func (rcv *InstaBotService) ReplyHTML(chatID int64, messageID int, text string) error {
	message := tgbotapi.NewMessage(chatID, text)
	message.ParseMode = tgbotapi.ModeHTML
	message.ReplyToMessageID = messageID
	_, err := rcv.bot.Send(message)
	return err
}

func (rcv *InstaBotService) ReplyAudio(chatID int64, messageID int, audio tgbotapi.AudioConfig) error {
	message := tgbotapi.NewMessage(chatID, "")
	message.ReplyToMessageID = messageID
	message.AllowSendingWithoutReply = true

	_, err := rcv.bot.Send(audio)
	return err
}

func (rcv *InstaBotService) ReplyVideo(chatID int64, messageID int, video tgbotapi.VideoConfig, caption string) error {
	message := tgbotapi.NewMessage(chatID, "")
	message.ReplyToMessageID = messageID
	message.AllowSendingWithoutReply = true
	message.Text = caption

	_, err := rcv.bot.Send(video)
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
) (pollID tgbotapi.Message, err error) {
	answer := tgbotapi.NewPoll(chatID, caption, options...)
	answer.IsAnonymous = isAnon
	rsp, err := rcv.bot.Send(answer)
	return rsp, err
}
