package listener

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/haski007/insta-bot/internal/bot/listener/transform"
	"github.com/haski007/insta-bot/internal/storage"
	"github.com/sashabaranov/go-openai"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (rcv *InstaBotService) msgGrokConversation(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	messageID := update.Message.MessageID
	username := update.Message.From.UserName
	userID := update.Message.From.ID

	if rcv.grokSrv == nil {
		rcv.SendError(chatID, ErrInternalServerError)
		rcv.NotifyCreator("[msgGrokConversation] Grok service is not configured (missing GROK_API_KEY)")
		return
	}

	nextPrompt := strings.TrimPrefix(update.Message.Text, "g!")

	history, err := rcv.storage.GetGrokConversation(&storage.GetConversationReq{
		Username: username,
		UserID:   userID,
		ChatID:   chatID,
	})

	// ---> Add system role applied to current chat
	var systemReplica *storage.Replica
	if !containsRole(history, openai.ChatMessageRoleSystem) {
		role, err := rcv.storage.GetSystemRole(chatID)
		if err != nil && !errors.Is(err, storage.ErrNotFound) {
			rcv.NotifyCreator(fmt.Sprintf("[msgGrokConversation] get system role: %s\n", err))
		}

		if role != "" {
			systemReplica = &storage.Replica{
				Role:    openai.ChatMessageRoleSystem,
				Content: role,
			}
			history = append(history, *systemReplica)
		}
	}

	grokMessages := transform.ReplicasToGPTMessagesHistory(history)

	grokMessages = append(grokMessages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: nextPrompt,
	})

	rsp, err := rcv.grokSrv.Conversation(context.Background(), grokMessages)
	if err != nil {
		rcv.log.WithError(err).Error("[msgGrokConversation] conversation")
		rcv.SendError(chatID, ErrInternalServerError)
		rcv.NotifyCreator(fmt.Sprintf("[msgGrokConversation] conversation: %s\n", err))
		return
	}

	replicasToSave := []storage.Replica{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: nextPrompt,
		},
		{
			Role:    openai.ChatMessageRoleAssistant,
			Content: rsp,
		},
	}

	if systemReplica != nil {
		replicasToSave = append(replicasToSave, *systemReplica)
	}

	if err := rcv.storage.PushGrokConversation(&storage.PushConversationReq{
		Username: username,
		UserID:   userID,
		ChatID:   chatID,
		Replicas: replicasToSave,
	}); err != nil {
		rcv.log.WithError(err).Error("[msgGrokConversation] push conversation")
		rcv.SendError(chatID, ErrInternalServerError)
		rcv.NotifyCreator(fmt.Sprintf("[msgGrokConversation] push conversation: %s\n", err))
		return
	}

	rcv.log.WithFields(map[string]interface{}{
		"from":   username,
		"prompt": nextPrompt,
	}).Info("new message to grok conversation")

	if err := rcv.Reply(chatID, messageID, rsp); err != nil {
		rcv.SendError(chatID, ErrInternalServerError)
		rcv.NotifyCreator(fmt.Sprintf("[msgGrokConversation] send message: %s\n", err))
		return
	}
}
