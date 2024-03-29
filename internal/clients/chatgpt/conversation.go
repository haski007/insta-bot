package chatgpt

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

func (srv *Service) Conversation(ctx context.Context, promptWithHistory []openai.ChatCompletionMessage) (string, error) {
	res, err := srv.ai.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:    srv.convGPTModel,
		Messages: promptWithHistory,
	})
	if err != nil {
		return "", fmt.Errorf("create chat completion: %w", err)
	}

	return res.Choices[0].Message.Content, nil
}
