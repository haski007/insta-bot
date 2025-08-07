package listener

import (
	"fmt"
	"strings"

	"github.com/haski007/insta-bot/internal/bot/publisher"
	"github.com/haski007/insta-bot/pkg/emoji"
	"github.com/haski007/insta-bot/pkg/safego"
	"github.com/sirupsen/logrus"
)

func (rcv *InstaBotService) StartPool() error {
	defer func() {
		if err := recover(); err != nil {
			rcv.log.WithError(fmt.Errorf("%s", err)).Error("[Pooling] panic")
			rcv.NotifyCreator(fmt.Sprintf("[Pooling] panic: %s", err))
			return
		}
	}()
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
				if err := rcv.sendStartInfo(update); err != nil {
					rcv.log.WithError(err).Println("[new chat member update] send start info")
				}
			}()

			continue
		}

		// if it's any type of media but the caption contains command /w
		if update.Message.Command() == "w" || strings.HasPrefix(update.Message.Caption, "/w") {
			go rcv.cmdWriteToChat(update)
			continue
		}

		go rcv.streamMessageToChats(update.Message)

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

			// PUBG addon ^)
			case command == "reg_pubg_players":
				go rcv.cmdRegPUBGPlayersHandler(update)
			case command == "purge_pubg_players":
				go rcv.cmdPurgePUBGPlayersHandler(update)
			case command == "lets_play_pubg":
				go rcv.cmdLetsPlayPUBGHandler(update)

			// Finals addon ^)
			case command == "reg_finals_players":
				go rcv.cmdRegFinalsPlayersHandler(update)
			case command == "purge_finals_players":
				go rcv.cmdPurgeFinalsPlayersHandler(update)
			case command == "lets_play_finals":
				go rcv.cmdLetsPlayFinalsHandler(update)

			// Users
			case command == "set_email":
				go rcv.cmdSetEmailHandler(update)

			case command == "set_system_role":
				go rcv.cmdSetSystemRoleHandler(update)
			case command == "drop_my_gpt":
				go rcv.cmdDropGPTConversationHandler(update)

			case command == "spam":
				go rcv.cmdSpam(update)

			case command == "sub_to_startup":
				go rcv.cmdSubToStartupHandler(update)
			case command == "unsub_to_startup":
				go rcv.cmdUnsubToStartupHandler(update)

			case command == "sum":
				safego.New(func() {
					rcv.cmdSum(update)
				}, func(pErr any) {
					rcv.log.WithError(fmt.Errorf("%s", pErr)).Error("[cmdSum] panic")
					rcv.NotifyCreator(fmt.Sprintf("%s cmdSum panic: %s", emoji.NoEntry, pErr))
					return
				})
			case command == "purge_history":
				go rcv.cmdPurgeHistory(update)

			case command == "stream_chat":
				go rcv.cmdStreamChat(update)
			case command == "stop_stream_chat":
				go rcv.cmdStopStreamChat(update)
			case command == "get_streams":
				go rcv.cmdGetStreamingChats(update)

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
			case strings.Contains(update.Message.Text, publisher.InstagramBaseUrl):
				logrus.Infof("Instagram post ignored: %s", update.Message.Text)
				// go rcv.msgInstagramTrigger(update)
				rcv.log.Infof("Ignore instagram post: %s due to broken downloader", update.Message.Text)
			case strings.Contains(update.Message.Text, publisher.TwitterBaseUrl), strings.Contains(update.Message.Text, publisher.TwitterOLDBaseUrl):
				go rcv.msgTwitterTrigger(update)
			case strings.Contains(update.Message.Text, publisher.InstagramStoriesBaseUrl):
				//go rcv.msgStoriesTrigger(update)
				rcv.log.Infof("Ignore stories: %s due to broken downloader", update.Message.Text)

			case strings.Contains(update.Message.Text, publisher.TikTokBaseUrl) ||
				strings.Contains(update.Message.Text, publisher.TikTokShareBaseUrl):
				//go rcv.msgTikTokTrigger(update)
				rcv.log.Infof("Ignore tiktok: %s due to broken downloader", update.Message.Text)

			case strings.Contains(update.Message.Text, publisher.YoutubeVideoBaseUrl):
				//go rcv.msgYoutubeTrigger(update)
				rcv.log.Infof("Ignore youtube: %s due to broken downloader", update.Message.Text)

			case len([]rune(update.Message.Text)) > 0 && []rune(update.Message.Text)[0] == '?' && len([]rune(update.Message.Text)) > 1:
				go rcv.msgChatGPTQuestion(update)
			case len([]rune(update.Message.Text)) > 0 && []rune(update.Message.Text)[0] == '!' && len([]rune(update.Message.Text)) > 1:
				go rcv.msgChatGTPConversation(update)
			case len([]rune(update.Message.Text)) > 0 && []rune(update.Message.Text)[0] == '~' && len([]rune(update.Message.Text)) > 1:
				go rcv.msgGPTextToSpeech(update)

			}

			// ---> save to history
			go rcv.msgSaveToHistory(update)
		}
	}

	logrus.Printf("Channel is closed")
	return nil
}
