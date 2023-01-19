package listener

import (
	"context"
	"fmt"

	"github.com/haski007/insta-bot/internal/bot"
	"github.com/haski007/insta-bot/internal/clients/google"
	"github.com/haski007/insta-bot/internal/clients/tiktokapi"
	"github.com/haski007/insta-bot/internal/clients/youtube"
	"github.com/haski007/insta-bot/internal/storage"
	"github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	CSGOContext = "csgo"
	PUBGContext = "pubg"
)

type InstaBotService struct {
	bot        *tgbotapi.BotAPI
	tiktokApi  *tiktokapi.TikTokClient
	youtubeApi *youtube.Client
	instapi    bot.InstApi
	calendar   google.Calendar
	updates    tgbotapi.UpdatesChannel
	storage    storage.Storage

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
	youtubeApi *youtube.Client,
	storage storage.Storage,
	calendarSrv google.Calendar,
) *InstaBotService {
	return &InstaBotService{
		bot:               botApi,
		creatorID:         creatorID,
		updates:           updatesChan,
		instapi:           instapi,
		ctx:               ctx,
		captionCharsLimit: captionCharsLimit,
		tiktokApi:         tiktokApi,
		youtubeApi:        youtubeApi,
		storage:           storage,
		calendar:          calendarSrv,
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
