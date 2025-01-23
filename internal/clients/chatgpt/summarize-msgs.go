package chatgpt

import (
	"context"
	"fmt"
	"strings"

	"github.com/sashabaranov/go-openai"
)

// SummarizeMessages handles messages larger than 4096 tokens.
func (srv *Service) SummarizeMessages(ctx context.Context, messages []string) (string, error) {
	// Assume maxTokens is the maximum number of tokens allowed in a single request.
	const maxTokens = 4096 * 2

	// Split messages into chunks that fit within the token limit.
	messageChunks := splitMessagesByTokens(messages, maxTokens)

	// Process each chunk and concatenate the results.
	var complMessages []openai.ChatCompletionMessage
	for i, chunk := range messageChunks {
		var prompt string
		if i == 0 {
			prompt += `Привіт GPT, я надаватиму тобі серію повідомлень частинами, кожна з яких містить частину поточної розмови. Ці повідомлення є сегментами більшої дискусії, і я надсилатиму їх послідовно. Після того, як я надішлю всі частини, я вкажу на кінець вводу, надіславши спеціальне фінальне повідомлення "End of conversation."
Коли ти отримаєш повідомлення "End of conversation", будь ласка, надай короткий але вичерпний підсумок всієї розмови. У своєму підсумку включи ключові обговорені моменти, будь-які досягнуті висновки та основні теми. Твій підсумок повинен охопити суть розмови та представити її чітко та лаконічно. Твоя відповідь має бути мовою, яка переважно використовувалась у переданих повідомленнях (українською).
Ось перша частина розмови:
`
		}
		prompt += strings.Join(chunk, "\n")

		complMessages = append(complMessages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: prompt,
		})
		promptWithHistory := openai.ChatCompletionRequest{
			Model:    srv.historyGPTModel,
			Messages: complMessages,
		}

		res, err := srv.ai.CreateChatCompletion(ctx, promptWithHistory)
		if err != nil {
			return "", fmt.Errorf("create chat completion: %w", err)
		}

		complMessages = append(complMessages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: res.Choices[0].Message.Content,
		})
	}

	complMessages = append(complMessages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: "End of conversation",
	})
	finalSummary, err := srv.ai.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:    srv.historyGPTModel,
		Messages: complMessages,
	})
	if err != nil {
		return "", fmt.Errorf("create chat completion: %w", err)
	}

	return finalSummary.Choices[0].Message.Content, nil
}

// Function to split messages into smaller chunks.
func splitMessagesByTokens(messages []string, limit int) [][]string {
	var chunks [][]string
	var currentChunk []string
	var currentLength int

	for _, msg := range messages {
		// Estimate the token count for the message.
		// This is a simplification; a more accurate method may be needed.
		tokenCount := len(strings.Fields(msg))

		if currentLength+tokenCount > limit && len(currentChunk) > 0 {
			chunks = append(chunks, currentChunk)
			currentChunk = nil
			currentLength = 0
		}

		currentChunk = append(currentChunk, msg)
		currentLength += tokenCount
	}

	if len(currentChunk) > 0 {
		chunks = append(chunks, currentChunk)
	}

	return chunks
}
