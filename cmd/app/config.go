package main

import (
	"time"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	// Google
	GoogleCredentialsPath string `env:"GOOGLE_CREDENTIALS_PATH"`

	// Youtube
	MaxQuality int `env:"MAX_QUALITY"`

	// OpenAI
	OpenAIAPIKey             string `env:"OPENAI_API_KEY"`
	OpenAIGPTModelForConv    string `env:"OPENAI_GPT_MODEL_FOR_CONV"`
	OpenAIGPTModelForHistory string `env:"OPENAI_GPT_MODEL_FOR_HISTORY"`

	// Grok (xAI)
	GrokAPIKey       string `env:"GROK_API_KEY"`
	GrokModelForConv string `env:"GROK_MODEL_FOR_CONV" envDefault:"grok-2-latest"`

	// OpenRouter (cheap models for extras, e.g. anglicism hints)
	OpenRouterAPIKey string `env:"OPENROUTER_API_KEY"`
	OpenRouterModel  string `env:"OPENROUTER_MODEL" envDefault:"google/gemma-2-9b-it"`
	// Max runes per message to send to anglicism LLM; longer messages are skipped (no API call).
	UkraineAnglicismMaxMessageRunes int `env:"UKRAINE_ANGLICISM_MAX_MESSAGE_RUNES" envDefault:"1000"`
	// Delivery mode for anglicism replies: "video" | "photo" | "plain". Default is video.
	// Reasoning: enemy bot iteratively adds filters; switch via env without redeploy.
	UkraineAnglicismDelivery string `env:"UKRAINE_ANGLICISM_DELIVERY" envDefault:"video"`

	// Redis
	RedisAddr               string        `env:"REDIS_ADDR"`
	RedisPass               string        `env:"REDIS_PASS"`
	RedisConversationTTL    time.Duration `env:"REDIS_CONVERSATION_TTL"`
	RedisHistoryMessagesTTL time.Duration `env:"REDIS_HISTORY_MESSAGES_TTL"`
	CaptionCharsLimit       int           `env:"CAPTION_CHARS_LIMIT"`

	// Telegram bot
	Token             string `env:"TELEGRAM_BOT_TOKEN"`
	UpdatesTimeoutSec int    `env:"TELEGRAM_BOT_UPDATES_TIMEOUT_SEC"`
	CreatorUserID     int64  `env:"TELEGRAM_BOT_CREATOR_USER_ID"`

	// Instloader
	InstloaderBaseURL string `env:"INSTLOADER_BASE_URL"`

	// ARC raiders
	ARCRAidersBaseURL string `env:"ARC_RAIDERS_BASE_URL"`
}

func Load(cfg *Config) error {
	if err := env.Parse(cfg); err != nil {
		return err
	}
	return nil
}
