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

func (rcv *InstaBotService) msgChatGTPConversation(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	messageID := update.Message.MessageID
	username := update.Message.From.UserName
	userID := update.Message.From.ID

	nextPrompt := strings.TrimPrefix(update.Message.Text, "!")

	history, err := rcv.storage.GetConversation(&storage.GetConversationReq{
		Username: username,
		UserID:   userID,
		ChatID:   chatID,
	})

	// ---> Add system role applied to current chat
	var systemReplica *storage.Replica
	if !containsRole(history, openai.ChatMessageRoleSystem) {
		role, err := rcv.storage.GetSystemRole(chatID)
		if err != nil && !errors.Is(err, storage.ErrNotFound) {
			rcv.NotifyCreator(fmt.Sprintf("[msgChatGTPConversation] get system role: %s\n", err))
		}

		if role != "" {
			systemReplica = &storage.Replica{
				Role:    openai.ChatMessageRoleSystem,
				Content: role,
			}
			history = append(history, *systemReplica)
		}
	}

	gptMessages := transform.ReplicasToGPTMessagesHistory(history)

	gptMessages = append(gptMessages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: nextPrompt,
	})

	// Get answer from GPT on prompt that depends on history of conversation
	rsp, err := rcv.gpt.Conversation(context.Background(), gptMessages)
	if err != nil {
		rcv.log.WithError(err).Error("[msgChatGPTQuestion] get short answer")
		rcv.SendError(chatID, ErrInternalServerError)
		rcv.NotifyCreator(fmt.Sprintf("[msgChatGPTQuestion] get short answer: %s\n", err))
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

	// Push to the redis history of conversation new prompt and response
	if systemReplica != nil {
		replicasToSave = append(replicasToSave, *systemReplica)
	}

	if err := rcv.storage.PushConversation(&storage.PushConversationReq{
		Username: username,
		UserID:   userID,
		ChatID:   chatID,
		Replicas: replicasToSave},
	); err != nil {
		rcv.log.WithError(err).Error("[msgChatGPTQuestion] push conversation")
		rcv.SendError(chatID, ErrInternalServerError)
		rcv.NotifyCreator(fmt.Sprintf("[msgChatGPTQuestion] push conversation: %s\n", err))
		return
	}

	rcv.log.WithFields(map[string]interface{}{
		"from":   username,
		"prompt": nextPrompt,
	}).Info("new message to chat gpt conversation")

	if err := rcv.Reply(chatID, messageID, rsp); err != nil {
		rcv.SendError(chatID, ErrInternalServerError)
		rcv.NotifyCreator(fmt.Sprintf("[msgChatGPTQuestion] send message: %s\n", err))
		return
	}
}

func containsRole(replicas []storage.Replica, targetRole string) bool {
	for _, r := range replicas {
		if r.Role == targetRole {
			return true
		}
	}
	return false
}
