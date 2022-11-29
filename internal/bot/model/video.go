package model

import (
	"time"

	"github.com/haski007/insta-bot/pkg/file"
)

type Video struct {
	Title        string `json:"title"`
	DownloadUrl  string `json:"download_url"`
	OriginalLink string `json:"original_link"`
	Author       string `json:"author"`

	// Additional fields for YouTube videos
	QualityLabel string        `json:"quality_label"`
	Views        int           `json:"views"`
	Duration     time.Duration `json:"duration"`
	CreatedAt    time.Time     `json:"created_at"`
}

func (v *Video) DownloadAsFile(path, filename string) (filePath string, err error) {
	return file.Download(v.DownloadUrl, path, filename)
}

func (v *Video) GetAuthor() string {
	return v.Author
}

func (v *Video) GetOriginalLink() string {
	return v.OriginalLink
}

func (v *Video) GetTitle() string {
	return v.Title
}

func (v *Video) GetDownloadUrl() string {
	return v.DownloadUrl
}
