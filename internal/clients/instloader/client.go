package instloader

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL *url.URL
}

func NewClient(baseURL *url.URL) *Client {
	return &Client{
		BaseURL: baseURL,
	}
}

// MediaItem is one slide of a post (single image/video or carousel entry).
type MediaItem struct {
	IsVideo  bool   `json:"is_video"`
	URL      string `json:"url"`
	VideoURL string `json:"video_url"`
}

type PostInfo struct {
	Shortcode string      `json:"shortcode"`
	IsVideo   bool        `json:"is_video"`
	URL       string      `json:"url"`
	VideoURL  string      `json:"video_url"`
	Caption   string      `json:"caption"`
	Owner     string      `json:"owner"`
	Likes     int         `json:"likes"`
	Comments  int         `json:"comments"`
	Timestamp string      `json:"timestamp"`
	Media     []MediaItem `json:"media"`
}

// MediaItems returns carousel slides when present, otherwise a single item
// built from the legacy top-level url/video_url fields.
func (p PostInfo) MediaItems() []MediaItem {
	if len(p.Media) > 0 {
		return p.Media
	}
	if p.VideoURL == "" && p.URL == "" {
		return nil
	}
	return []MediaItem{{
		IsVideo:  p.IsVideo,
		URL:      p.URL,
		VideoURL: p.VideoURL,
	}}
}

func (c *Client) GetPostInfo(shortcode string) (PostInfo, error) {
	reqURL := *c.BaseURL
	reqURL.Path = "/media"
	reqURL.RawQuery = url.Values{
		"shortcode": {shortcode},
	}.Encode()
	return c.getMediaInfo(reqURL)
}

// GetStoryInfo fetches a single Instagram story by numeric media id.
func (c *Client) GetStoryInfo(mediaID, username string) (PostInfo, error) {
	reqURL := *c.BaseURL
	reqURL.Path = "/story"
	values := url.Values{
		"media_id": {mediaID},
	}
	if username != "" {
		values.Set("username", username)
	}
	reqURL.RawQuery = values.Encode()
	return c.getMediaInfo(reqURL)
}

func (c *Client) getMediaInfo(reqURL url.URL) (PostInfo, error) {
	req, err := http.NewRequest(http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return PostInfo{}, fmt.Errorf("create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return PostInfo{}, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PostInfo{}, fmt.Errorf("read body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return PostInfo{}, fmt.Errorf("instloader http %d: %s", resp.StatusCode, string(body))
	}

	var postInfo PostInfo
	if err := json.Unmarshal(body, &postInfo); err != nil {
		return PostInfo{}, fmt.Errorf("unmarshal post info: %w", err)
	}

	return postInfo, nil
}
