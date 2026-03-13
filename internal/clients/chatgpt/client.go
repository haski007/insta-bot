package chatgpt

import (
	"errors"

	"github.com/sashabaranov/go-openai"
)

var (
	ErrInvalidGPTModel = errors.New("invalid gpt model")
)

type Service struct {
	ai              *openai.Client
	convGPTModel    string
	historyGPTModel string
	apiKey          string
}

func NewService(
	ai *openai.Client,
	convGPTModel,
	historyGPTModel,
	apiKey string,
) (*Service, error) {
	if err := validateGPTModels(convGPTModel, historyGPTModel); err != nil {
		return nil, err
	}

	return &Service{
		ai:              ai,
		convGPTModel:    convGPTModel,
		historyGPTModel: historyGPTModel,
		apiKey:          apiKey,
	}, nil
}

func validateGPTModels(model ...string) error {
	for _, m := range model {
		if m == "" {
			return ErrInvalidGPTModel
		}
	}

	return nil
}
