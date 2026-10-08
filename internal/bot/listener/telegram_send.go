package listener

import (
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/haski007/insta-bot/internal/metrics"
)

func telegramSendKind(c tgbotapi.Chattable) string {
	switch c.(type) {
	case tgbotapi.MessageConfig:
		return "message"
	case tgbotapi.PhotoConfig:
		return "photo"
	case tgbotapi.VideoConfig:
		return "video"
	case tgbotapi.AudioConfig:
		return "audio"
	case tgbotapi.DocumentConfig:
		return "document"
	case tgbotapi.VoiceConfig:
		return "voice"
	case tgbotapi.VideoNoteConfig:
		return "video_note"
	case tgbotapi.ForwardConfig:
		return "forward"
	case tgbotapi.DeleteMessageConfig:
		return "delete"
	case tgbotapi.InvoiceConfig:
		return "invoice"
	case tgbotapi.MediaGroupConfig:
		return "media_group"
	default:
		return "other"
	}
}

func (rcv *InstaBotService) sendAPI(c tgbotapi.Chattable) (tgbotapi.Message, error) {
	start := time.Now()
	msg, err := rcv.bot.Send(c)
	metrics.ObserveTelegramSend(telegramSendKind(c), start, err)
	return msg, err
}

func (rcv *InstaBotService) sendMediaGroupAPI(cfg tgbotapi.MediaGroupConfig) ([]tgbotapi.Message, error) {
	start := time.Now()
	msgs, err := rcv.bot.SendMediaGroup(cfg)
	metrics.ObserveTelegramSend("media_group", start, err)
	return msgs, err
}
