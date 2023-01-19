package main

import (
	"fmt"
	"os"

	"github.com/haski007/insta-bot/pkg/factory"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Clients struct {
		Instapi    factory.HTTPClient `yaml:"instapi"`
		YoutubeApi YouTubeConfig      `yaml:"youtube_api"`
		Redis      RedisClient        `yaml:"redis"`
		Google     GoogleConfig       `yaml:"google"`
	}
	TelegramBot factory.TelegramBotCfg `yaml:"telegram_bot"`

	CaptionCharsLimit int `yaml:"caption_chars_limit"`
}

type GoogleConfig struct {
	CredentialsPath string `yaml:"credentials_path"`
}

type RedisClient struct {
	Addr string `yaml:"addr"`
}

type YouTubeConfig struct {
	MaxQuality int `yaml:"max_quality"`
}

func Load(configFile string, cfg interface{}) error {
	fileData, err := os.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("can't read config file: %w", err)
	}

	fileData = []byte(os.ExpandEnv(string(fileData)))

	if err = yaml.Unmarshal(fileData, cfg); err != nil {
		return fmt.Errorf("can't unmarshal config data: %w", err)
	}

	if v, ok := cfg.(interface{ Validate() error }); ok {
		if err = v.Validate(); err != nil {
			return fmt.Errorf("invalid config: %w", err)
		}
	}

	return nil
}
