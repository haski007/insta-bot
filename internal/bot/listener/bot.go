package listener

import (
	"context"
	"fmt"
	"time"

	"github.com/haski007/insta-bot/internal/clients/instapi"

	"github.com/haski007/insta-bot/internal/clients/chatgpt"

	"github.com/haski007/insta-bot/internal/clients/google"
	"github.com/haski007/insta-bot/internal/clients/tiktokapi"
	"github.com/haski007/insta-bot/internal/clients/youtube"
	"github.com/haski007/insta-bot/internal/storage"
	"github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	CSGOContext   = "csgo"
	PUBGContext   = "pubg"
	FinalsContext = "finals"

	PollContext  = "poll"
	UsersContext = "users"
)

type InstaBotService struct {
	bot        *tgbotapi.BotAPI
	tiktokApi  *tiktokapi.TikTokClient
	youtubeApi *youtube.Client
	instapi    *instapi.Api
	calendar   google.Calendar
	updates    tgbotapi.UpdatesChannel
	storage    storage.Storage
	gpt        *chatgpt.Service

	creatorID         int64
	captionCharsLimit int
	log               logrus.FieldLogger

	ctx context.Context
}

func NewInstaBotService(
	ctx context.Context,
	botApi *tgbotapi.BotAPI,
	instapi *instapi.Api,
	creatorID int64,
	updatesChan tgbotapi.UpdatesChannel,
	captionCharsLimit int,
	tiktokApi *tiktokapi.TikTokClient,
	youtubeApi *youtube.Client,
	storage storage.Storage,
	calendarSrv google.Calendar,
	gpt *chatgpt.Service,
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
		gpt:               gpt,
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

func (rcv *InstaBotService) RunAfterFuncsPolls() error {
	polls, err := rcv.storage.GetAllPolls()
	if err != nil {
		return fmt.Errorf("get all polls from redis err: %w", err)
	}

	for _, poll := range polls {
		if poll.Time.After(time.Now()) {
			time.AfterFunc(poll.Time.Sub(time.Now()), func() {
				if err := rcv.SendMessage(poll.ChatID, "Here we go guys! "+poll.MeetLink); err != nil {
					rcv.log.WithError(err).Println("[afterFunc] send message")
					return
				}
			})
			rcv.log.WithField("starts_at", poll.Time.String()).
				Info("added afterFunc to send google meet")
		}

	}

	return nil
}
