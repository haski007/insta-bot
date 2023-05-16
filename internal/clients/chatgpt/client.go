package chatgpt

import "github.com/sashabaranov/go-openai"

type Service struct {
	ai *openai.Client
}

func NewService(ai *openai.Client) *Service {
	return &Service{
		ai: ai,
	}
}
