package chatgpt

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

func (srv *Service) Conversation(ctx context.Context, promptWithHistory []openai.ChatCompletionMessage) (string, error) {
	res, err := srv.ai.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:    openai.GPT3Dot5Turbo,
		Messages: promptWithHistory,
	})
	if err != nil {
		return "", fmt.Errorf("create chat completion: %w", err)
	}

	return res.Choices[0].Message.Content, nil
}

func (srv *Service) ConversationGPT4(ctx context.Context, promptWithHistory []openai.ChatCompletionMessage) (string, error) {
	res, err := srv.ai.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:    openai.GPT4,
		Messages: promptWithHistory,
	})
	if err != nil {
		return "", fmt.Errorf("create chat completion: %w", err)
	}

	return res.Choices[0].Message.Content, nil
}
