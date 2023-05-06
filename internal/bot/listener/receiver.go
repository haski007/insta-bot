package listener

import (
	"fmt"
	"strings"

	"github.com/haski007/insta-bot/internal/bot/publisher"
	"github.com/haski007/insta-bot/pkg/emoji"
	"github.com/sirupsen/logrus"
)

func (rcv *InstaBotService) StartPool() error {
	me, err := rcv.bot.GetMe()
	if err != nil {
		_ = rcv.NotifyCreator(fmt.Sprintf("[bot GetMe] err: %s", err))
		return err
	}

	for update := range rcv.updates {
		if update.EditedMessage != nil || update.Poll != nil {
			continue
		}

		if update.PollAnswer != nil {
			go rcv.triggerPollAnswer(update)
			continue
		}

		// Check if someone added bot to chat
		if update.MyChatMember != nil &&
			update.MyChatMember.NewChatMember.User.ID == me.ID {
			go func() {
				if err := rcv.sendStartInfo(update.MyChatMember.Chat.ID); err != nil {
					rcv.log.WithError(err).Println("[new chat member update] send start info")
				}
			}()
			continue
		}

		// ---> Commands
		if update.Message != nil && update.Message.IsCommand() {
			command := update.Message.Command()
			switch {
			case command == "test":
				go rcv.cmdTestHandler(update)
			case command == "help":
				go rcv.cmdStartHandler(update)

			case command == "set_quality":
				go rcv.cmdSetQualityHandler(update)

			case command == "list_players":
				go rcv.cmdListPlayersHandler(update)

			// CSGO addon ^)
			case command == "reg_csgo_players":
				go rcv.cmdRegCSGOPlayersHandler(update)
			case command == "purge_csgo_players":
				go rcv.cmdPurgeCSGOPlayersHandler(update)

			case command == "lets_play":
				go rcv.cmdLetsPlayHandler(update)

			// CSGO addon ^)
			case command == "reg_pubg_players":
				go rcv.cmdRegPUBGPlayersHandler(update)
			case command == "purge_pubg_players":
				go rcv.cmdPurgePUBGPlayersHandler(update)
			case command == "lets_play_pubg":
				go rcv.cmdLetsPlayPUBGHandler(update)

			// Users
			case command == "set_email":
				go rcv.cmdSetEmailHandler(update)

			default:
				go func() {
					if err := rcv.SendMessage(
						update.Message.Chat.ID,
						"Such command does not exist! "+emoji.NoEntry,
					); err != nil {
						logrus.WithError(err).Printf("send message to chat: %d", update.Message.Chat.ID)
					}
				}()
			}
		}

		// Parse messages
		if update.Message != nil && !update.Message.IsCommand() {
			switch {
			case strings.Contains(update.Message.Text, "https://www.instagram.com/"):
				//go rcv.msgMediaTrigger(update)
				rcv.log.Infof("Ignore instagram post: %s due to broken downloader", update.Message.Text)
			case strings.Contains(update.Message.Text, "https://instagram.com/stories"):
				//go rcv.msgStoriesTrigger(update)
				rcv.log.Infof("Ignore stories: %s due to broken downloader", update.Message.Text)

			case strings.Contains(update.Message.Text, publisher.TikTokBaseUrl) ||
				strings.Contains(update.Message.Text, publisher.TikTokShareBaseUrl):
				//go rcv.msgTikTokTrigger(update)
				rcv.log.Infof("Ignore tiktok: %s due to broken downloader", update.Message.Text)

			case strings.Contains(update.Message.Text, publisher.YoutubeVideoBaseUrl):
				go rcv.msgYoutubeTrigger(update)

			case len([]rune(update.Message.Text)) > 0 && []rune(update.Message.Text)[0] == '?':
				go rcv.msgChatGPTQuestion(update)
			case len([]rune(update.Message.Text)) > 0 && []rune(update.Message.Text)[0] == '!':
				go rcv.msgChatGTPConversation(update)
			}

		}
	}

	logrus.Printf("Channel is closed")
	return nil
}
