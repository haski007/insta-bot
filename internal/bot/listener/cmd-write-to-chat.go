package listener

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/haski007/insta-bot/pkg/emoji"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (rcv *InstaBotService) cmdWriteToChat(update tgbotapi.Update) {
	if !rcv.IsCreator(update.Message.From.ID) {
		rcv.SendMessage(update.Message.Chat.ID, "You are not allowed to use this command"+emoji.NoEntry)
		return
	}

	chatID := update.Message.Chat.ID
	messageID := update.Message.MessageID

	text := update.Message.Text
	if text == "" {
		text = update.Message.Caption
	}

	// Split the message text into arguments
	args := rcv.parseCommandArgs(text)
	if len(args) < 1 { // Changed to 1 since message text is optional when sending media
		msg := tgbotapi.NewMessage(chatID, "Usage: /w <target_chat_id> [message]")
		msg.ReplyToMessageID = messageID
		rcv.bot.Send(msg)
		return
	}

	// Parse target chat ID from the first argument
	targetChatID, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		msg := tgbotapi.NewMessage(chatID, "Invalid chat ID. Please provide a valid number.")
		msg.ReplyToMessageID = messageID
		rcv.bot.Send(msg)
		return
	}

	// Handle media from the original message
	switch {
	case update.Message.Photo != nil && len(update.Message.Photo) > 0:
		photoMsg := tgbotapi.NewPhoto(targetChatID, tgbotapi.FileID(update.Message.Photo[len(update.Message.Photo)-1].FileID))
		if len(args) > 1 {
			photoMsg.Caption = strings.Join(args[1:], " ")
		}
		_, err = rcv.bot.Send(photoMsg)

	case update.Message.Video != nil:
		videoMsg := tgbotapi.NewVideo(targetChatID, tgbotapi.FileID(update.Message.Video.FileID))
		if len(args) > 1 {
			videoMsg.Caption = strings.Join(args[1:], " ")
		}
		_, err = rcv.bot.Send(videoMsg)

	case update.Message.Document != nil:
		docMsg := tgbotapi.NewDocument(targetChatID, tgbotapi.FileID(update.Message.Document.FileID))
		if len(args) > 1 {
			docMsg.Caption = strings.Join(args[1:], " ")
		}
		_, err = rcv.bot.Send(docMsg)

	case update.Message.Voice != nil:
		voiceMsg := tgbotapi.NewVoice(targetChatID, tgbotapi.FileID(update.Message.Voice.FileID))
		if len(args) > 1 {
			voiceMsg.Caption = strings.Join(args[1:], " ")
		}
		_, err = rcv.bot.Send(voiceMsg)

	case update.Message.VideoNote != nil:
		videoNoteMsg := tgbotapi.NewVideoNote(targetChatID, 0, tgbotapi.FileID(update.Message.VideoNote.FileID))
		_, err = rcv.bot.Send(videoNoteMsg)

	default:
		// If no media in original message, check for reply message
		if update.Message.ReplyToMessage != nil {
			replyMsg := update.Message.ReplyToMessage

			// Handle different types of media
			switch {
			case replyMsg.Photo != nil && len(replyMsg.Photo) > 0:
				photoMsg := tgbotapi.NewPhoto(targetChatID, tgbotapi.FileID(replyMsg.Photo[len(replyMsg.Photo)-1].FileID))
				if len(args) > 1 {
					photoMsg.Caption = strings.Join(args[1:], " ")
				}
				_, err = rcv.bot.Send(photoMsg)

			case replyMsg.Video != nil:
				videoMsg := tgbotapi.NewVideo(targetChatID, tgbotapi.FileID(replyMsg.Video.FileID))
				if len(args) > 1 {
					videoMsg.Caption = strings.Join(args[1:], " ")
				}
				_, err = rcv.bot.Send(videoMsg)

			case replyMsg.Document != nil:
				docMsg := tgbotapi.NewDocument(targetChatID, tgbotapi.FileID(replyMsg.Document.FileID))
				if len(args) > 1 {
					docMsg.Caption = strings.Join(args[1:], " ")
				}
				_, err = rcv.bot.Send(docMsg)

			case replyMsg.Voice != nil:
				voiceMsg := tgbotapi.NewVoice(targetChatID, tgbotapi.FileID(replyMsg.Voice.FileID))
				if len(args) > 1 {
					voiceMsg.Caption = strings.Join(args[1:], " ")
				}
				_, err = rcv.bot.Send(voiceMsg)

			case replyMsg.VideoNote != nil:
				videoNoteMsg := tgbotapi.NewVideoNote(targetChatID, 0, tgbotapi.FileID(replyMsg.VideoNote.FileID))
				_, err = rcv.bot.Send(videoNoteMsg)

			default:
				// If no media, just forward the message text if it exists
				if len(args) > 1 {
					msg := tgbotapi.NewMessage(targetChatID, strings.Join(args[1:], " "))
					_, err = rcv.bot.Send(msg)
				} else {
					msg := tgbotapi.NewMessage(chatID, "No message or media to send.")
					msg.ReplyToMessageID = messageID
					rcv.bot.Send(msg)
					return
				}
			}

			if err != nil {
				errorMsg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Failed to send media: %v", err))
				errorMsg.ReplyToMessageID = messageID
				rcv.bot.Send(errorMsg)
				return
			}

		} else {
			// No media and no reply message, just send text if provided
			if len(args) > 1 {
				msg := tgbotapi.NewMessage(targetChatID, strings.Join(args[1:], " "))
				_, err = rcv.bot.Send(msg)
				if err != nil {
					errorMsg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Failed to send message: %v", err))
					errorMsg.ReplyToMessageID = messageID
					rcv.bot.Send(errorMsg)
					return
				}
			} else {
				msg := tgbotapi.NewMessage(chatID, "No message to send.")
				msg.ReplyToMessageID = messageID
				rcv.bot.Send(msg)
				return
			}
		}
	}

	// Send confirmation to the command sender
	confirmMsg := tgbotapi.NewMessage(chatID, "Message sent successfully! "+emoji.Check)
	confirmMsg.ReplyToMessageID = messageID
	rcv.bot.Send(confirmMsg)
}