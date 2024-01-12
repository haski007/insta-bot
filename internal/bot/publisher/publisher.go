package publisher

const (
	InstagramBaseUrl        = "https://www.instagram.com/"
	InstagramStoriesBaseUrl = "https://instagram.com/"

	TwitterBaseUrl    = "https://x.com/"
	TwitterOLDBaseUrl = "https://twitter.com/"

	VXTwitterBaseUrl = "https://vxtwitter.com/"

	TikTokBaseUrl      = "https://www.tiktok.com/"
	TikTokShareBaseUrl = "https://vm.tiktok.com/"

	YoutubeBaseUrl      = "https://www.youtube.com"
	YoutubeVideoBaseUrl = "https://www.youtube.com/watch"
)

type Publisher struct {
	authToken        string
	username         string
	password         string
	verificationCode string
}
