package bot

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
