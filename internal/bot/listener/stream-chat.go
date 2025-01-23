package listener

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/haski007/insta-bot/pkg/emoji"
	"github.com/haski007/pretty"
	"github.com/sirupsen/logrus"
)

func (rcv *InstaBotService) cmdStreamChat(update tgbotapi.Update) {
	if !rcv.IsCreator(update.Message.From.ID) {
		rcv.SendMessage(update.Message.Chat.ID, "You are not allowed to use this command" + emoji.NoEntry)
		return
	}

	var chatIDToStream int64

	args := rcv.parseCommandArgs(update.Message.Text)
	if len(args) < 1 {
		rcv.SendMessage(update.Message.Chat.ID, "Usage: /stream_chat <chat_id>")
		return
	}

	chatIDToStream, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		rcv.SendMessage(update.Message.Chat.ID, "Invalid chat ID")
		return
	}
	if value, ok := rcv.streamChats.Load(chatIDToStream); ok {
		arr := value.([]int64)
		arr = append(arr, update.Message.Chat.ID)
		rcv.streamChats.Store(chatIDToStream, arr)
	} else {
		rcv.streamChats.Store(chatIDToStream, []int64{update.Message.Chat.ID})
	}
	
	rcv.SendMessage(update.Message.Chat.ID, "Chat will be streamed to "+strconv.FormatInt(update.Message.Chat.ID, 10) + emoji.Check)
}

func (rcv *InstaBotService) cmdStopStreamChat(update tgbotapi.Update) {
	if !rcv.IsCreator(update.Message.From.ID) {
		rcv.SendMessage(update.Message.Chat.ID, "You are not allowed to use this command" + emoji.NoEntry)
		return
	}
	var chatIDToStop int64

	args := rcv.parseCommandArgs(update.Message.Text)
	if len(args) < 1 {
		rcv.SendMessage(update.Message.Chat.ID, "Usage: /stop_stream_chat <chat_id>")
		return
	}

	chatIDToStop, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		rcv.SendMessage(update.Message.Chat.ID, "Invalid chat ID")
		return
	}

	rcv.streamChats.Delete(chatIDToStop)
	rcv.SendMessage(update.Message.Chat.ID, "Chat will be stopped from streaming to "+strconv.FormatInt(chatIDToStop, 10))
}

func (rcv *InstaBotService) streamMessageToChats(message *tgbotapi.Message) {
	if value, ok := rcv.streamChats.Load(message.Chat.ID); ok {
		chatsToStream := value.([]int64)
		for _, chatID := range chatsToStream {
			rcv.bot.Send(tgbotapi.NewForward(chatID, message.Chat.ID, message.MessageID))
			logrus.WithField("from_chat_id", chatID).WithField("from_chat_title", message.Chat.Title).Info("message was forwarded")
		}
	}
}

func (rcv *InstaBotService) cmdGetStreamingChats(update tgbotapi.Update) {
	var result []int64
	rcv.streamChats.Range(func(key, value any) bool {
		chatID := key.(int64)
		chatsToWrite := value.([]int64)
		for _, chw := range chatsToWrite {
			if chw == update.Message.Chat.ID {
				result = append(result, chatID)
			}
		}
		
		return true
	})

	rcv.SendMessage(update.Message.Chat.ID, fmt.Sprintf("Streaming chats: %s", pretty.String(result)))
}
