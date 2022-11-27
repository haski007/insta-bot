package listener

import (
	"context"
	"fmt"

	"github.com/haski007/insta-bot/internal/bot"
	"github.com/haski007/insta-bot/internal/clients/tiktokapi"
	"github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type InstaBotService struct {
	bot       *tgbotapi.BotAPI
	tiktokApi *tiktokapi.TikTokClient
	instapi   bot.InstApi
	updates   tgbotapi.UpdatesChannel

	creatorID         int64
	captionCharsLimit int
	log               logrus.FieldLogger

	ctx context.Context
}

func NewInstaBotService(
	ctx context.Context,
	botApi *tgbotapi.BotAPI,
	instapi bot.InstApi,
	creatorID int64,
	updatesChan tgbotapi.UpdatesChannel,
	captionCharsLimit int,
	tiktokApi *tiktokapi.TikTokClient,
) *InstaBotService {
	return &InstaBotService{
		bot:               botApi,
		creatorID:         creatorID,
		updates:           updatesChan,
		instapi:           instapi,
		ctx:               ctx,
		captionCharsLimit: captionCharsLimit,
		tiktokApi:         tiktokApi,
	}
}

func (rcv *InstaBotService) SetLogger(logger logrus.FieldLogger) *InstaBotService {
	rcv.log = logger.WithField("handler", "rcv")
	return rcv
}

func (rcv *InstaBotService) NotifyCreator(message string) error {
	if err := rcv.SendMessage(rcv.creatorID, message); err != nil {
		return fmt.Errorf("notify creator err: %w", err)
	}
	return nil
}

func (rcv *InstaBotService) StopPool(_ context.Context) error {
	if err := rcv.NotifyCreator("Bot got signal so it's shutting down..."); err != nil {
		return fmt.Errorf("notify creator err: %w", err)
	}
	rcv.bot.StopReceivingUpdates()
	return nil
}
