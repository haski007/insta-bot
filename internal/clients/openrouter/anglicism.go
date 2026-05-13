package openrouter

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/sashabaranov/go-openai"
)

type AnglicismResult struct {
	HasAnglicism bool
	Rewritten    string
}

type anglicismJSON struct {
	HasAnglicism bool   `json:"has_anglicism"`
	Rewritten    string `json:"rewritten"`
}

const anglicismSystemPrompt = `Ти допомагаєш україномовним чатам. Твоє завдання — знайти в тексті повідомлення англіцизми: зайві запозичення з англійської там, де є природні українські відповідники (наприклад розмовні "крінж", "імба", "фічерити", "колаба" тощо). Не вважай англіцизмами: власні назви, бренди, нікнейми, URL, технічні терміни без доброго українського аналога, латиницю в коді або командах.

Якщо зайвих англіцизмів немає — поверни JSON: {"has_anglicism":false,"rewritten":""}

Якщо є — поверни JSON: {"has_anglicism":true,"rewritten":"..."} де rewritten — той самий текст повідомлення, але кожен виправлений фрагмент СТРОГО в форматі {{український відповідник}}(було: англіцизм) без змін у написанні дужок і слова «було:». Інший текст без HTML/markdown. Відповідай лише одним JSON-об'єктом без пояснень.`

func (s *Service) AnalyzeAnglicisms(ctx context.Context, userMessage string) (*AnglicismResult, error) {
	if s == nil || s.ai == nil {
		return nil, fmt.Errorf("openrouter: service not configured")
	}
	res, err := s.ai.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: s.model,
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleSystem, Content: anglicismSystemPrompt},
			{Role: openai.ChatMessageRoleUser, Content: "Проаналізуй це повідомлення і поверни JSON.\n\nТекст:\n" + userMessage},
		},
		Temperature: 0.15,
		MaxTokens:   800,
		ResponseFormat: &openai.ChatCompletionResponseFormat{
			Type: openai.ChatCompletionResponseFormatTypeJSONObject,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("openrouter completion: %w", err)
	}
	if len(res.Choices) == 0 {
		return nil, fmt.Errorf("openrouter: empty choices")
	}
	raw := strings.TrimSpace(res.Choices[0].Message.Content)
	raw = stripJSONFences(raw)

	var parsed anglicismJSON
	if err := json.Unmarshal([]byte(raw), &parsed); err != nil {
		return nil, fmt.Errorf("openrouter parse json: %w (body: %q)", err, truncate(raw, 400))
	}
	return &AnglicismResult{HasAnglicism: parsed.HasAnglicism, Rewritten: strings.TrimSpace(parsed.Rewritten)}, nil
}

func stripJSONFences(s string) string {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "```") {
		s = strings.TrimPrefix(s, "```")
		s = strings.TrimSpace(s)
		s = strings.TrimPrefix(s, "json")
		s = strings.TrimSpace(s)
		if i := strings.Index(s, "```"); i >= 0 {
			s = strings.TrimSpace(s[:i])
		}
	}
	return s
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max] + "…"
}
