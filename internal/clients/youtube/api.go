package youtube

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/haski007/insta-bot/internal/bot/model"
	"github.com/haski007/insta-bot/pkg/text"
	"github.com/kkdai/youtube/v2"
)

type Client struct {
	conn       *youtube.Client
	maxQuality int
}

func New(maxQuality int) *Client {
	return &Client{
		conn:       new(youtube.Client),
		maxQuality: maxQuality,
	}
}

func (rcv *Client) SetMaxQuality(value int) {
	rcv.maxQuality = value
}

func (rcv *Client) GetMaxQuality() int {
	return rcv.maxQuality
}

func (rcv *Client) GetVideoByUrl(url string, qualityLimit int) (*model.Video, error) {
	video, err := rcv.conn.GetVideo(url)
	if err != nil {
		return nil, fmt.Errorf("get video err: %w", err)
	}

	if len(video.Formats) == 0 {
		return nil, model.ErrNoVideoFound
	}

	var maxQuality, maxIdx int
	for i, format := range video.Formats {
		if format.QualityLabel == "" {
			continue
		}
		quality, err := strconv.Atoi(strings.ReplaceAll(format.QualityLabel, "p", ""))
		if err != nil {
			return nil, fmt.Errorf("atoi to quality label err: %w", err)
		}

		if quality > maxQuality && quality <= rcv.GetMaxQuality() {
			maxQuality = quality
			maxIdx = i
		}
	}

	result := video.Formats[maxIdx]

	var downloadUrl string
	if result.URL != "" {
		downloadUrl = result.URL
	}

	if downloadUrl == "" {
		return nil, model.ErrNoVideoFound
	}

	downloadUrl = text.DecodeUrl(downloadUrl)

	return &model.Video{
		Title:        video.Title,
		DownloadUrl:  downloadUrl,
		OriginalLink: url,
		Author:       video.Author,

		Views:        video.Views,
		Duration:     video.Duration,
		QualityLabel: result.QualityLabel,
		CreatedAt:    video.PublishDate,
	}, nil
}
