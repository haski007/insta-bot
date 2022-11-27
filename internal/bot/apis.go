package bot

import (
	"context"

	"github.com/haski007/insta-bot/internal/clients/instapi"
)

type InstApi interface {
	Login(ctx context.Context) error

	GetMediaPkFromUrl(ctx context.Context, url string) (int, error)
	GetMediaInfoFromPk(ctx context.Context, pk int) (*instapi.Media, error)
	GetMediaInfoFromURL(ctx context.Context, url string) (*instapi.Media, error)

	GetStoryPkFromUrl(ctx context.Context, url string) (string, error)
	GetStoryInfoFromPk(ctx context.Context, pk string) (*instapi.Story, error)
	GetStoryInfoFromUrl(ctx context.Context, url string) (*instapi.Story, error)
}

type TikTokApi interface {
	GetVideoDataFromUrl(url string) (TikTokVideo, error)
}

type TikTokVideo interface {
	DownloadAsFile(path, filename string) (filePath string, err error)
	GetAuthor() string
	GetOriginalLink() string
	GetDownloadUrl() string
	GetTitle() string
}
