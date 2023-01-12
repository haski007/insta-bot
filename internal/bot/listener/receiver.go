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
		if update.EditedMessage != nil || update.Poll != nil || update.PollAnswer != nil {
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

			// CSGO addon ^)
			case command == "reg_csgo_players":
				go rcv.cmdRegCSGOPlayersHandler(update)
			case command == "purge_csgo_players":
				go rcv.cmdPurgeCSGOPlayersHandler(update)
			case command == "list_players":
				go rcv.cmdListPlayersHandler(update)
			case command == "lets_play":
				go rcv.cmdLetsPlayHandler(update)

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
				go rcv.msgMediaTrigger(update)
			case strings.Contains(update.Message.Text, "https://instagram.com/stories"):
				go rcv.msgStoriesTrigger(update)

			case strings.Contains(update.Message.Text, publisher.TikTokBaseUrl) ||
				strings.Contains(update.Message.Text, publisher.TikTokShareBaseUrl):
				go rcv.msgTikTokTrigger(update)

			case strings.Contains(update.Message.Text, publisher.YoutubeVideoBaseUrl):
				go rcv.msgYoutubeTrigger(update)
			}

		}
	}

	logrus.Printf("Channel is closed")
	return nil
}
