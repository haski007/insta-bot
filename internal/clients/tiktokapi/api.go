package tiktokapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/haski007/insta-bot/internal/bot"
	"github.com/haski007/insta-bot/pkg/file"
)

const (
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36"
	// Telegram Bot API rejects files over 50MB.
	telegramBotFileLimit = 50 * 1024 * 1024
	defaultAPIBase       = "https://www.tikwm.com/api/"
	tikwmOrigin          = "https://www.tikwm.com"
)

type TikTokClient struct {
	apiBase string
	http    *http.Client
}

func New() *TikTokClient {
	return &TikTokClient{
		apiBase: defaultAPIBase,
		http:    &http.Client{Timeout: 30 * time.Second},
	}
}

// Media is a TikTok video or photo slideshow ready to send to Telegram.
type Media struct {
	ID        string
	Title     string
	Author    string
	SourceURL string
	VideoURL  string
	Cover     string
	Images    []string
}

func (m *Media) GetAuthor() string       { return m.Author }
func (m *Media) GetTitle() string        { return m.Title }
func (m *Media) GetOriginalLink() string { return m.SourceURL }
func (m *Media) GetDownloadUrl() string  { return m.VideoURL }

func (m *Media) DownloadAsFile(path, filename string) (string, error) {
	if m.VideoURL == "" {
		return "", fmt.Errorf("no video url")
	}
	return file.Download(m.VideoURL, path, filename)
}

var _ bot.TikTokVideo = (*Media)(nil)

type apiResponse struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data *apiData `json:"data"`
}

type apiData struct {
	ID     string    `json:"id"`
	Title  string    `json:"title"`
	Play   string    `json:"play"`
	WMPlay string    `json:"wmplay"`
	HDPlay string    `json:"hdplay"`
	Size   int64     `json:"size"`
	HDSize int64     `json:"hd_size"`
	Cover  string    `json:"cover"`
	Images []string  `json:"images"`
	Author apiAuthor `json:"author"`
}

type apiAuthor struct {
	UniqueID string `json:"unique_id"`
	Nickname string `json:"nickname"`
}

func (rcv *TikTokClient) GetVideoDataFromUrl(rawURL string) (bot.TikTokVideo, error) {
	media, err := rcv.GetMedia(rawURL)
	if err != nil {
		return nil, err
	}
	return media, nil
}

func (rcv *TikTokClient) GetMedia(rawURL string) (*Media, error) {
	rawURL = strings.TrimSpace(rawURL)
	if rawURL == "" {
		return nil, fmt.Errorf("empty tiktok url")
	}

	payload, err := rcv.fetch(rawURL)
	if err != nil {
		return nil, err
	}
	if payload.Code != 0 {
		msg := strings.TrimSpace(payload.Msg)
		if msg == "" {
			msg = "tiktok api error"
		}
		return nil, fmt.Errorf("tiktok api: %s", msg)
	}
	if payload.Data == nil {
		return nil, fmt.Errorf("tiktok api returned no data")
	}

	return mediaFromAPI(rawURL, payload.Data), nil
}

func (rcv *TikTokClient) fetch(tiktokURL string) (*apiResponse, error) {
	endpoint, err := url.Parse(rcv.apiBase)
	if err != nil {
		return nil, fmt.Errorf("parse tiktok api base: %w", err)
	}
	query := endpoint.Query()
	query.Set("url", tiktokURL)
	query.Set("hd", "1")
	endpoint.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("new tiktok api request: %w", err)
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "application/json")

	resp, err := rcv.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("tiktok api request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.LimitReader(resp.Body, 8<<20))
	if err != nil {
		return nil, fmt.Errorf("read tiktok api response: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("tiktok api http %d: %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}

	var payload apiResponse
	if err := json.Unmarshal(body, &payload); err != nil {
		return nil, fmt.Errorf("decode tiktok api response: %w", err)
	}
	return &payload, nil
}

func mediaFromAPI(sourceURL string, data *apiData) *Media {
	author := strings.TrimSpace(data.Author.UniqueID)
	if author == "" {
		author = strings.TrimSpace(data.Author.Nickname)
	}

	images := make([]string, 0, len(data.Images))
	for _, img := range data.Images {
		if abs := absolutize(img); abs != "" {
			images = append(images, abs)
		}
	}

	return &Media{
		ID:        strings.TrimSpace(data.ID),
		Title:     strings.TrimSpace(data.Title),
		Author:    author,
		SourceURL: sourceURL,
		VideoURL:  pickPlayURL(data),
		Cover:     absolutize(data.Cover),
		Images:    images,
	}
}

func pickPlayURL(data *apiData) string {
	hd := absolutize(data.HDPlay)
	play := absolutize(data.Play)
	wm := absolutize(data.WMPlay)

	if hd != "" && (data.HDSize == 0 || data.HDSize <= telegramBotFileLimit) {
		return hd
	}
	if play != "" && (data.Size == 0 || data.Size <= telegramBotFileLimit) {
		return play
	}
	if hd != "" {
		return hd
	}
	if play != "" {
		return play
	}
	return wm
}

func absolutize(raw string) string {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return ""
	}
	parsed, err := url.Parse(raw)
	if err != nil {
		return raw
	}
	if parsed.IsAbs() {
		return raw
	}
	base, err := url.Parse(tikwmOrigin)
	if err != nil {
		return raw
	}
	return base.ResolveReference(parsed).String()
}
