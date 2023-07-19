package chatgpt

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

const (
	QuestionDelimiter = "```"
)

func (srv *Service) GetShortAnswer(ctx context.Context, question string) (string, error) {
	prompt := fmt.Sprintf(`You  will be provided with a question 
delimited with triple backticks,
You should answer this question by a shortest way 
using language of the question.
Find relevant information and then answer the question
based on the relevant information. If there is no
such an relevant information do not generate random info
just add to the answer that your are not sure.
Question: %s", question)`, QuestionDelimiter+question+QuestionDelimiter)

	res, err := srv.ai.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	})
	if err != nil {
		return "", fmt.Errorf("create chat completion: %w", err)
	}

	return res.Choices[0].Message.Content, nil
}

func (srv *Service) GetShortAnswerGPT4(ctx context.Context, question string) (string, error) {
	prompt := fmt.Sprintf(`You  will be provided with a question 
delimited with triple backticks,
You should answer this question by a shortest way 
using language of the question.
Find relevant information and then answer the question
based on the relevant information. If there is no
such an relevant information do not generate random info
just add to the answer that your are not sure.
Question: %s", question)`, QuestionDelimiter+question+QuestionDelimiter)

	res, err := srv.ai.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT4,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	})
	if err != nil {
		return "", fmt.Errorf("create chat completion: %w", err)
	}

	return res.Choices[0].Message.Content, nil
}
