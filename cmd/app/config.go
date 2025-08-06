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
}

func Load(cfg *Config) error {
	if err := env.Parse(cfg); err != nil {
		return err
	}
	return nil
}
