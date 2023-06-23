package model

import (
	"time"

	"github.com/haski007/insta-bot/pkg/file"
)

type Author struct {
	Identifier Identifier `json:"identifier"`
	Name       string     `json:"name"`
	URL        string     `json:"URL"`
}

type Identifier struct {
	Value string `json:"value"`
}

func (a *Author) GetName() string {
	return a.Identifier.Value
}

func (a *Author) GetProfileURL() string {
	return a.URL
}

type Video struct {
	Name        string    `json:"name"`
	DownloadUrl string    `json:"contentUrl"`
	Description string    `json:"description"`
	Caption     string    `json:"caption"`
	UploadDate  time.Time `json:"uploadDate"`

	// Deprecated
	Title string `json:"title"`
	// Deprecated
	OriginalLink string `json:"original_link"`
	// Deprecated
	Author string `json:"author"`
	// Deprecated
	Views int `json:"views"`
	// Deprecated
	Duration time.Duration `json:"duration"`
	// Deprecated
	QualityLabel string `json:"quality_label"`
	// Deprecated
	CreatedAt time.Time `json:"created_at"`
}

func (v *Video) DownloadAsFile(path, filename string) (filePath string, err error) {
	return file.Download(v.DownloadUrl, path, filename)
}

func (v *Video) Getname() string {
	return v.Name
}

func (v *Video) GetAuthor() string {
	return v.Author
}

func (v *Video) GetTitle() string {
	return v.Title
}

func (v *Video) GetDownloadUrl() string {
	return v.DownloadUrl
}

type Image struct {
	DownloadUrl string `json:"url"`
	Caption     string `json:"caption"`
}

func (v *Image) DownloadAsFile(path, filename string) (filePath string, err error) {
	return file.Download(v.DownloadUrl, path, filename)
}

func (v *Image) GetDownloadUrl() string {
	return v.DownloadUrl
}

func (v *Image) GetCaption() string {
	return v.Caption
}
