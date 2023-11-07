package chatgpt

import (
	"errors"

	"github.com/sashabaranov/go-openai"
)

var (
	ErrInvalidGPTModel = errors.New("invalid gpt model")
)

type Service struct {
	ai       *openai.Client
	gptModel string
	apiKey   string
}

func NewService(ai *openai.Client, gptModel, apiKey string) (*Service, error) {
	if err := validateGPTModel(gptModel); err != nil {
		return nil, err
	}

	return &Service{
		ai:       ai,
		gptModel: gptModel,
		apiKey:   apiKey,
	}, nil
}

func validateGPTModel(model string) error {
	if model != openai.GPT3Dot5Turbo && model != openai.GPT4 {
		return ErrInvalidGPTModel
	}

	return nil
}
