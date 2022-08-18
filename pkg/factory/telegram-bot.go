package factory

type TelegramBotCfg struct {
	Token             string `json:"token" yaml:"token"`
	CreatorUserID     int64  `json:"creator_user_id" yaml:"creator_user_id"`
	UpdatesTimeoutSec int    `json:"updates_timeout_sec" yaml:"updates_timeout_sec"`
}
