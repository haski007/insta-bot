package listener

import (
	"fmt"
	"strings"

	"github.com/haski007/insta-bot/internal/bot/publisher"
	"github.com/haski007/insta-bot/internal/metrics"
	"github.com/haski007/insta-bot/pkg/emoji"
	"github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
		metrics.IncUpdate(updateKind(update))

		if update.EditedMessage != nil || update.Poll != nil {
			continue
		}

		if update.Message != nil {
			go rcv.reactPoopIfNeeded(update.Message)
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
		if update.Message != nil && (update.Message.Command() == "w" || strings.HasPrefix(update.Message.Caption, "/w")) {
			rcv.goCommand("w", update, rcv.cmdWriteToChat)
			continue
		}

		if update.Message != nil {
			go rcv.streamMessageToChats(update.Message)
		}

		// ---> Commands
		if update.Message != nil && update.Message.IsCommand() {
			command := update.Message.Command()
			switch command {
			case "test":
				rcv.goCommand(command, update, rcv.cmdTestHandler)
			case "help":
				rcv.goCommand(command, update, rcv.cmdStartHandler)

			case "set_quality":
				rcv.goCommand(command, update, rcv.cmdSetQualityHandler)

			case "list_players":
				rcv.goCommand(command, update, rcv.cmdListPlayersHandler)

			case "reg_csgo_players":
				rcv.goCommand(command, update, rcv.cmdRegCSGOPlayersHandler)
			case "purge_csgo_players":
				rcv.goCommand(command, update, rcv.cmdPurgeCSGOPlayersHandler)

			case "lets_play":
				rcv.goCommand(command, update, rcv.cmdLetsPlayHandler)

			case "reg_pubg_players":
				rcv.goCommand(command, update, rcv.cmdRegPUBGPlayersHandler)
			case "purge_pubg_players":
				rcv.goCommand(command, update, rcv.cmdPurgePUBGPlayersHandler)
			case "lets_play_pubg":
				rcv.goCommand(command, update, rcv.cmdLetsPlayPUBGHandler)

			case "reg_finals_players":
				rcv.goCommand(command, update, rcv.cmdRegFinalsPlayersHandler)
			case "purge_finals_players":
				rcv.goCommand(command, update, rcv.cmdPurgeFinalsPlayersHandler)
			case "lets_play_finals":
				rcv.goCommand(command, update, rcv.cmdLetsPlayFinalsHandler)

			case "set_email":
				rcv.goCommand(command, update, rcv.cmdSetEmailHandler)

			case "set_system_role":
				rcv.goCommand(command, update, rcv.cmdSetSystemRoleHandler)
			case "drop_my_gpt":
				rcv.goCommand(command, update, rcv.cmdDropGPTConversationHandler)
			case "drop_my_grok":
				rcv.goCommand(command, update, rcv.cmdDropGrokConversationHandler)

			case "spam":
				rcv.goCommand(command, update, rcv.cmdSpam)

			case "sub_to_startup":
				rcv.goCommand(command, update, rcv.cmdSubToStartupHandler)
			case "unsub_to_startup":
				rcv.goCommand(command, update, rcv.cmdUnsubToStartupHandler)

			case "disable_loader":
				rcv.goCommand(command, update, rcv.cmdDisableLoaderHandler)
			case "enable_loader":
				rcv.goCommand(command, update, rcv.cmdEnableLoaderHandler)

			case "sub_arc_events":
				rcv.goCommand(command, update, rcv.cmdSubARCEventHandler)
			case "unsub_arc_events":
				rcv.goCommand(command, update, rcv.cmdUnsubARCEventHandler)
			case "arc":
				rcv.goCommand(command, update, rcv.cmdListArcEventsHandler)

			case "ukraine_for_ukrainians":
				rcv.goCommand(command, update, rcv.cmdUkraineForUkrainiansSub)
			case "unsub_ukraine_for_ukrainians":
				rcv.goCommand(command, update, rcv.cmdUkraineForUkrainiansUnsub)

			case "ignore":
				rcv.goCommand(command, update, rcv.cmdUkraineAnglicismIgnore)
			case "unignore":
				rcv.goCommand(command, update, rcv.cmdUkraineAnglicismUnignore)

			case "sum":
				rcv.goCommandSafe(command, update, rcv.cmdSum)
			case "purge_history":
				rcv.goCommand(command, update, rcv.cmdPurgeHistory)

			case "stream_chat":
				rcv.goCommand(command, update, rcv.cmdStreamChat)
			case "stop_stream_chat":
				rcv.goCommand(command, update, rcv.cmdStopStreamChat)
			case "get_streams":
				rcv.goCommand(command, update, rcv.cmdGetStreamingChats)

			case "fuck":
				rcv.goCommand(command, update, rcv.cmdFuck)
			case "unfuck":
				rcv.goCommand(command, update, rcv.cmdUnfuck)

			default:
				rcv.goCommand("unknown", update, func(update tgbotapi.Update) {
					if err := rcv.SendMessage(
						update.Message.Chat.ID,
						"Such command does not exist! "+emoji.NoEntry,
					); err != nil {
						logrus.WithError(err).Printf("send message to chat: %d", update.Message.Chat.ID)
					}
				})
			}
		}

		// Parse messages
		if update.Message != nil && !update.Message.IsCommand() {
			switch {
			case strings.Contains(update.Message.Text, "instagram.com"):
				loaderEnabled, err := rcv.storage.IsChatLoaderEnabled(update.Message.Chat.ID)
				if err != nil {
					rcv.log.WithError(err).Error("IsChatLoaderEnabled")
				}
				if !loaderEnabled {
					metrics.IncRoute("instagram_disabled")
					rcv.log.Infof("Ignore instagram url: %s due to loader disabled", update.Message.Text)
					break
				}
				igURL := exprFindURL.FindString(update.Message.Text)
				switch {
				case strings.Contains(igURL, "/stories/"):
					rcv.goMessage("stories", update, rcv.msgStoriesTrigger)
				case strings.Contains(igURL, "/p/"), strings.Contains(igURL, "/reel/"):
					rcv.goMessage("instagram", update, rcv.msgInstagramTrigger)
				default:
					metrics.IncRoute("instagram_unsupported")
					rcv.log.Infof("Ignore unsupported instagram url: %s", igURL)
				}
			case strings.Contains(update.Message.Text, publisher.TwitterBaseUrl), strings.Contains(update.Message.Text, publisher.TwitterOLDBaseUrl):
				metrics.IncRoute("twitter_ignored")
				rcv.log.Infof("Ignore twitter: %s", update.Message.Text)

			case isTikTokURL(update.Message.Text):
				loaderEnabled, err := rcv.storage.IsChatLoaderEnabled(update.Message.Chat.ID)
				if err != nil {
					rcv.log.WithError(err).Error("IsChatLoaderEnabled")
				}
				if !loaderEnabled {
					metrics.IncRoute("tiktok_disabled")
					rcv.log.Infof("Ignore tiktok url: %s due to loader disabled", update.Message.Text)
					break
				}
				rcv.goMessage("tiktok", update, rcv.msgTikTokTrigger)

			case strings.Contains(update.Message.Text, publisher.YoutubeVideoBaseUrl):
				metrics.IncRoute("youtube_ignored")
				rcv.log.Infof("Ignore youtube: %s due to broken downloader", update.Message.Text)

			case strings.HasPrefix(update.Message.Text, "?") && len(update.Message.Text) > 1:
				rcv.goMessage("gpt_question", update, rcv.msgChatGPTQuestion)
			case strings.HasPrefix(update.Message.Text, "!") && len(update.Message.Text) > 1:
				rcv.goMessage("gpt_conversation", update, rcv.msgChatGTPConversation)
			case strings.HasPrefix(update.Message.Text, "~") && len(update.Message.Text) > 1:
				rcv.goMessage("tts", update, rcv.msgGPTextToSpeech)

			case strings.HasPrefix(update.Message.Text, "g?") && len(update.Message.Text) > 2:
				rcv.goMessage("grok_question", update, rcv.msgGrokQuestion)
			case strings.HasPrefix(update.Message.Text, "g!") && len(update.Message.Text) > 2:
				rcv.goMessage("grok_conversation", update, rcv.msgGrokConversation)
			default:
				rcv.goMessage("anglicism", update, rcv.msgUkraineAnglicismIfNeeded)
			}
			// ---> save to history
			go rcv.msgSaveToHistory(update)
		}
	}

	logrus.Printf("Channel is closed")
	return nil
}
