package publisher

import (
	"github.com/haski007/insta-bot/internal/bot"
	"github.com/haski007/insta-bot/internal/clients/instapi"
)

const (
	InstagramBaseUrl        = "https://www.instagram.com/"
	InstagramStoriesBaseUrl = "https://instagram.com/"

	TikTokBaseUrl      = "https://www.tiktok.com/"
	TikTokShareBaseUrl = "https://vm.tiktok.com/"
)

type Publisher struct {
	apiCli *instapi.APIClient

	authToken        string
	username         string
	password         string
	verificationCode string
}

func New(cli *instapi.APIClient, username, password, verificationCode string) bot.InstApi {
	return &Publisher{
		apiCli: cli,
		// TODO: make empty after debug
		authToken:        "54655610031%3A8ss25xkfZkRCmz%3A7%3AAYeZBJ5RbWTpmlGcA5kHYvoDK-cOMoczMnCEpWCRTw",
		username:         username,
		password:         password,
		verificationCode: verificationCode,
	}
}
