package youtube

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/haski007/insta-bot/internal/metrics"
	"github.com/haski007/insta-bot/pkg/file"
)

const (
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36"

	defaultDownloadBase = "https://p.savenow.to/ajax/download.php"
	defaultPollInterval = time.Second
	defaultPollTimeout  = 90 * time.Second
)

type Client struct {
	downloadBase string
	http         *http.Client
	maxQuality   int
	pollInterval time.Duration
	pollTimeout  time.Duration
}

func New(maxQuality int) *Client {
	if maxQuality <= 0 {
		maxQuality = 720
	}
	return &Client{
		downloadBase: defaultDownloadBase,
		http:         &http.Client{Timeout: 30 * time.Second},
		maxQuality:   maxQuality,
		pollInterval: defaultPollInterval,
		pollTimeout:  defaultPollTimeout,
	}
}

func (rcv *Client) SetMaxQuality(value int) {
	if value > 0 {
		rcv.maxQuality = value
	}
}

func (rcv *Client) GetMaxQuality() int {
	return rcv.maxQuality
}

// Media is a YouTube Short ready to send to Telegram.
type Media struct {
	ID        string
	Title     string
	Author    string
	SourceURL string
	VideoURL  string
	Thumbnail string
	Quality   string
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

type startResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Success     bool   `json:"success"`
	Message     string `json:"message"`
	ProgressURL string `json:"progress_url"`
	Thumbnail   string `json:"thumbnail_url"`
	Info        *struct {
		Title string `json:"title"`
		Image string `json:"image"`
	} `json:"info"`
}

type progressResponse struct {
	Progress    int    `json:"progress"`
	Success     any    `json:"success"`
	Text        string `json:"text"`
	Title       string `json:"title"`
	DownloadURL string `json:"download_url"`
	Thumbnail   string `json:"thumbnail_url"`
	Message     string `json:"message"`
}

func (rcv *Client) GetMedia(rawURL string) (media *Media, err error) {
	start := time.Now()
	defer func() {
		metrics.ObserveExternal("youtube", "get_media", start, err)
	}()

	rawURL = strings.TrimSpace(rawURL)
	if rawURL == "" {
		return nil, fmt.Errorf("empty youtube url")
	}

	started, err := rcv.startDownload(rawURL)
	if err != nil {
		return nil, err
	}
	if started.ProgressURL == "" {
		msg := strings.TrimSpace(started.Message)
		if msg == "" {
			msg = "youtube api returned no progress url"
		}
		return nil, fmt.Errorf("youtube api: %s", msg)
	}

	done, err := rcv.pollProgress(started.ProgressURL)
	if err != nil {
		return nil, err
	}
	if strings.TrimSpace(done.DownloadURL) == "" {
		msg := strings.TrimSpace(done.Text)
		if msg == "" {
			msg = strings.TrimSpace(done.Message)
		}
		if msg == "" {
			msg = "youtube api returned no download url"
		}
		return nil, fmt.Errorf("youtube api: %s", msg)
	}

	title := strings.TrimSpace(done.Title)
	if title == "" {
		title = strings.TrimSpace(started.Title)
	}
	if title == "" && started.Info != nil {
		title = strings.TrimSpace(started.Info.Title)
	}

	thumb := strings.TrimSpace(done.Thumbnail)
	if thumb == "" {
		thumb = strings.TrimSpace(started.Thumbnail)
	}
	if thumb == "" && started.Info != nil {
		thumb = strings.TrimSpace(started.Info.Image)
	}

	return &Media{
		ID:        strings.TrimSpace(started.ID),
		Title:     title,
		SourceURL: rawURL,
		VideoURL:  strings.TrimSpace(done.DownloadURL),
		Thumbnail: thumb,
		Quality:   formatQuality(rcv.maxQuality),
	}, nil
}

func (rcv *Client) startDownload(videoURL string) (*startResponse, error) {
	endpoint, err := url.Parse(rcv.downloadBase)
	if err != nil {
		return nil, fmt.Errorf("parse youtube api base: %w", err)
	}
	query := endpoint.Query()
	query.Set("url", videoURL)
	query.Set("format", formatQuality(rcv.maxQuality))
	endpoint.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("new youtube api request: %w", err)
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "application/json")

	resp, err := rcv.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("youtube api request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.LimitReader(resp.Body, 8<<20))
	if err != nil {
		return nil, fmt.Errorf("read youtube api response: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("youtube api http %d: %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}

	var payload startResponse
	if err := json.Unmarshal(body, &payload); err != nil {
		return nil, fmt.Errorf("decode youtube api response: %w", err)
	}
	return &payload, nil
}

func (rcv *Client) pollProgress(progressURL string) (*progressResponse, error) {
	deadline := time.Now().Add(rcv.pollTimeout)
	interval := rcv.pollInterval
	if interval <= 0 {
		interval = defaultPollInterval
	}

	for {
		payload, err := rcv.fetchProgress(progressURL)
		if err != nil {
			return nil, err
		}
		if payload.Progress >= 1000 || truthy(payload.Success) {
			return payload, nil
		}
		if time.Now().After(deadline) {
			msg := strings.TrimSpace(payload.Text)
			if msg == "" {
				msg = "timed out waiting for youtube download"
			}
			return nil, fmt.Errorf("youtube api: %s", msg)
		}
		time.Sleep(interval)
	}
}

func (rcv *Client) fetchProgress(progressURL string) (*progressResponse, error) {
	req, err := http.NewRequest(http.MethodGet, progressURL, nil)
	if err != nil {
		return nil, fmt.Errorf("new youtube progress request: %w", err)
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "application/json")

	resp, err := rcv.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("youtube progress request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.LimitReader(resp.Body, 2<<20))
	if err != nil {
		return nil, fmt.Errorf("read youtube progress response: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("youtube progress http %d: %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}

	var payload progressResponse
	if err := json.Unmarshal(body, &payload); err != nil {
		return nil, fmt.Errorf("decode youtube progress response: %w", err)
	}
	return &payload, nil
}

func formatQuality(maxQuality int) string {
	switch {
	case maxQuality <= 360:
		return "360"
	case maxQuality <= 480:
		return "480"
	case maxQuality <= 720:
		return "720"
	default:
		return "1080"
	}
}

func truthy(v any) bool {
	switch t := v.(type) {
	case bool:
		return t
	case float64:
		return t != 0
	case string:
		s := strings.ToLower(strings.TrimSpace(t))
		return s == "1" || s == "true" || s == "yes"
	case json.Number:
		n, err := t.Int64()
		return err == nil && n != 0
	default:
		if v == nil {
			return false
		}
		s := strings.TrimSpace(fmt.Sprint(v))
		if s == "" {
			return false
		}
		n, err := strconv.Atoi(s)
		if err == nil {
			return n != 0
		}
		return strings.EqualFold(s, "true")
	}
}
