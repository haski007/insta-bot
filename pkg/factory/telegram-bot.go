package factory

import "time"

type TelegramBotCfg struct {
	Token          string        `json:"token" yaml:"token" env:"BOT_TOKEN"`
	CreatorUserID  int64         `json:"creator_user_id" yaml:"creator_user_id" env:"BOT_CREATOR_USER_ID"`
	UpdatesTimeout time.Duration `json:"updates_timeout" yaml:"updates_timeout" env:"BOT_UPDATES_TIMEOUT"`
}
