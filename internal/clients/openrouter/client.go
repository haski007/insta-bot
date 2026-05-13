package openrouter

import (
	"errors"

	"github.com/sashabaranov/go-openai"
)

const APIBaseURL = "https://openrouter.ai/api/v1"

var ErrInvalidModel = errors.New("openrouter: model is empty")

type Service struct {
	ai    *openai.Client
	model string
}

func NewService(apiKey, model string) (*Service, error) {
	if apiKey == "" {
		return nil, nil
	}
	if model == "" {
		return nil, ErrInvalidModel
	}
	cfg := openai.DefaultConfig(apiKey)
	cfg.BaseURL = APIBaseURL
	ai := openai.NewClientWithConfig(cfg)
	return &Service{ai: ai, model: model}, nil
}
