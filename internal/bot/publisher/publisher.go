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

	YoutubeBaseUrl      = "https://www.youtube.com"
	YoutubeVideoBaseUrl = "https://www.youtube.com/watch"
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
		username:         username,
		password:         password,
		verificationCode: verificationCode,
	}
}
