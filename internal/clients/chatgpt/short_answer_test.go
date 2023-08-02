package chatgpt

import (
	"context"
	"os"
	"testing"

	"github.com/sashabaranov/go-openai"
	"github.com/stretchr/testify/assert"
)

func TestService_GetShortAnswer(t *testing.T) {
	key := os.Getenv("OPENAI_API_KEY")
	if key == "" {
		t.Skip("OPENAI_API_KEY is not set")
	}

	ai := openai.NewClient(key)
	srv, err := NewService(ai, "gpt3-dot-5-turbo")
	if err != nil {
		t.Error(err)
	}

	answer, err := srv.GetShortAnswer(context.Background(), "What is the capital of the USA?")
	if err != nil {
		t.Error(err)
	}
	assert.NotEmpty(t, answer)
}
