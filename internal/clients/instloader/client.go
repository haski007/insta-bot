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

type PostInfo struct {
	Shortcode string `json:"shortcode"`
	IsVideo   bool   `json:"is_video"`
	URL       string `json:"url"`
	VideoURL  string `json:"video_url"`
	Caption   string `json:"caption"`
	Owner     string `json:"owner"`
	Likes     int    `json:"likes"`
	Comments  int    `json:"comments"`
	Timestamp string `json:"timestamp"`
}

func (c *Client) GetPostInfo(shortcode string) (PostInfo, error) {
	req := http.Request{
		Method: http.MethodGet,
		URL:    c.BaseURL,
	}

	req.URL.Path = "/media"
	req.URL.RawQuery = url.Values{
		"shortcode": {shortcode},
	}.Encode()

	resp, err := http.DefaultClient.Do(&req)
	if err != nil {
		return PostInfo{}, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PostInfo{}, fmt.Errorf("read body: %w", err)
	}

	var postInfo PostInfo
	if err := json.Unmarshal(body, &postInfo); err != nil {
		return PostInfo{}, fmt.Errorf("unmarshal post info: %w", err)
	}

	return postInfo, nil
}