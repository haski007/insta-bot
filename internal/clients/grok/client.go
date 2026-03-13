package grok

import (
	"errors"

	"github.com/sashabaranov/go-openai"
)

const (
	APIBaseURL = "https://api.x.ai/v1"
)

var (
	ErrInvalidModel = errors.New("invalid grok model")
)

type Service struct {
	ai       *openai.Client
	convModel string
}

func NewService(apiKey, convModel string) (*Service, error) {
	if convModel == "" {
		return nil, ErrInvalidModel
	}

	cfg := openai.DefaultConfig(apiKey)
	cfg.BaseURL = APIBaseURL
	ai := openai.NewClientWithConfig(cfg)

	return &Service{
		ai:        ai,
		convModel: convModel,
	}, nil
}
