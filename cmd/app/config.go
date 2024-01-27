package main

import (
	"time"

	"github.com/haski007/insta-bot/pkg/factory"
)

type Config struct {
	Clients struct {
		YoutubeApi YouTubeConfig `yaml:"youtube_api"`
		Redis      RedisClient   `yaml:"redis"`
		Google     GoogleConfig  `yaml:"google"`
		OpenAI     OpenAIConfig  `yaml:"openai"`
	}
	TelegramBot factory.TelegramBotCfg `yaml:"telegram_bot"`

	CaptionCharsLimit int `yaml:"caption_chars_limit"`
}

type GoogleConfig struct {
	Credentials string `yaml:"credentials_path" env:"GOOGLE_CREDENTIALS"`
}

type OpenAIConfig struct {
	ApiKey             string `yaml:"api_key" env:"OPENAI_API_KEY"`
	GPTModelForConv    string `yaml:"gpt_model_for_conv" env:"OPENAI_GPT_MODEL_FOR_CONV"`
	GPTModelForHistory string `yaml:"gpt_model_for_history" env:"OPENAI_GPT_MODEL_FOR_HISTORY"`
}

type RedisClient struct {
	Addr               string        `yaml:"addr" env:"REDIS_ADDR"`
	Pass               string        `yaml:"pass" env:"REDIS_PASS"`
	ConversationTTL    time.Duration `yaml:"conversation_ttl_min" env:"REDIS_CONVERSATION_TTL"`
	HistoryMessagesTTL time.Duration `yaml:"history_messages_ttl_hours" env:"REDIS_HISTORY_MESSAGES_TTL"`
}

type YouTubeConfig struct {
	MaxQuality int `yaml:"max_quality"`
}
