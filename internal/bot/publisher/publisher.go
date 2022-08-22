package publisher

import "github.com/haski007/insta-bot/internal/clients/instapi"

const (
	InstagramBaseUrl        = "https://www.instagram.com/"
	InstagramStoriesBaseUrl = "https://instagram.com/"
)

type Publisher struct {
	apiCli *instapi.APIClient

	authToken string
	username  string
	password  string
}

func New(cli *instapi.APIClient, username, password string) *Publisher {
	return &Publisher{
		apiCli: cli,
		// TODO: make empty after debug
		authToken: "",
		username:  username,
		password:  password,
	}
}
