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
		if update.EditedMessage != nil {
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
		if update.Message.IsCommand() {
			command := update.Message.CommandWithAt()
			switch {
			case command == "test":
				go rcv.cmdTestHandler(update)

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
		if update.Message != nil {
			switch {
			case strings.Contains(update.Message.Text, "https://www.instagram.com/"):
				go rcv.msgMediaTrigger(update)
			case strings.Contains(update.Message.Text, "https://instagram.com/stories"):
				go rcv.msgStoriesTrigger(update)

			case strings.Contains(update.Message.Text, publisher.TikTokBaseUrl) ||
				strings.Contains(update.Message.Text, publisher.TikTokShareBaseUrl):
				go rcv.msgTikTokTrigger(update)

			}

		}
	}

	logrus.Printf("Channel is closed")
	return nil
}
