package model

import "github.com/haski007/insta-bot/pkg/file"

type TikTokVideo struct {
	Title        string `json:"title"`
	DownloadUrl  string `json:"download_url"`
	OriginalLink string `json:"original_link"`
	Author       string `json:"author"`
}

func (v *TikTokVideo) DownloadAsFile(path, filename string) (filePath string, err error) {
	return file.Download(v.DownloadUrl, path, filename)
}

func (v *TikTokVideo) GetAuthor() string {
	return v.Author
}

func (v *TikTokVideo) GetOriginalLink() string {
	return v.OriginalLink
}

func (v *TikTokVideo) GetTitle() string {
	return v.Title
}

func (v *TikTokVideo) GetDownloadUrl() string {
	return v.DownloadUrl
}
