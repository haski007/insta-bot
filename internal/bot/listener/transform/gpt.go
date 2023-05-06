package transform

import (
	"github.com/haski007/insta-bot/internal/storage"
	"github.com/sashabaranov/go-openai"
)

func ReplicasToGPTMessagesHistory(replicas []storage.Replica) (messages []openai.ChatCompletionMessage) {
	for _, replica := range replicas {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    replica.Role,
			Content: replica.Content,
			Name:    replica.Name,
		})
	}

	return
}
