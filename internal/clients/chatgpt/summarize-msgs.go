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
	const maxTokens = 4096

	// Split messages into chunks that fit within the token limit.
	messageChunks := splitMessagesByTokens(messages, maxTokens)

	// Process each chunk and concatenate the results.
	var complMessages []openai.ChatCompletionMessage
	for i, chunk := range messageChunks {
		var prompt string
		if i == 0 {
			prompt += `Hello GPT, I will be providing you with a series of messages in chunks, each containing part of an ongoing conversation. These messages are segments of a larger discussion, and I'll send them to you sequentially. After I have sent all the chunks, I will indicate the end of the input by sending a specific final message saying "End of conversation."
Once you receive the "End of conversation" message, please provide a comprehensive but short summary of the entire conversation. In your summary, include the key points discussed, any conclusions reached, and the main themes or topics covered. Your summary should capture the essence of the conversation and present it in a clear, concise manner. Your answer should be on language that mostly used in passed messages (Ukrainian).
Here's the first chunk of the conversation:
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
