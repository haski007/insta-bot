package tiktokapi

const (
	userAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36"
)

type TikTokClient struct{}

func New() *TikTokClient {
	return &TikTokClient{}
}
