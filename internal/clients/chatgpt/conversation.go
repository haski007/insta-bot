package chatgpt

import (
	"context"
	"fmt"
	"time"

	"github.com/haski007/insta-bot/internal/metrics"
	"github.com/sashabaranov/go-openai"
)

func (srv *Service) Conversation(ctx context.Context, promptWithHistory []openai.ChatCompletionMessage) (answer string, err error) {
	start := time.Now()
	defer func() {
		metrics.ObserveLLM("openai", "conversation", start, err)
	}()

	res, err := srv.ai.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:    srv.convGPTModel,
		Messages: promptWithHistory,
	})
	if err != nil {
		return "", fmt.Errorf("create chat completion: %w", err)
	}

	return res.Choices[0].Message.Content, nil
}
