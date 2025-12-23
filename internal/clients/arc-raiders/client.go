package arcraiders

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	eventsTimersPath = "/api/arc-raiders/event-timers"
)

type Client struct {
	BaseURL *url.URL
}

func NewClient(baseURL *url.URL) *Client {
	return &Client{
		BaseURL: baseURL,
	}
}

type TimeRange struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type EventTimer struct {
	Game        string      `json:"game"`
	Name        string      `json:"name"`
	Map         string      `json:"map"`
	Icon        string      `json:"icon"`
	Description string      `json:"description"`
	Days        []string    `json:"days"`
	Times       []TimeRange `json:"times"`
}

type EventTimersResponse struct {
	Data []EventTimer `json:"data"`
}

type GetARCEventsOptions struct {
	Map  string
	Name string
}

func (c *Client) GetARCEvents(opts *GetARCEventsOptions) ([]EventTimer, error) {
	req := http.Request{
		Method: http.MethodGet,
		URL:    c.BaseURL,
	}

	req.URL.Path = eventsTimersPath

	query := url.Values{}
	if opts != nil {
		if opts.Map != "" {
			query.Set("map", opts.Map)
		}
		if opts.Name != "" {
			query.Set("name", opts.Name)
		}
	}
	req.URL.RawQuery = query.Encode()

	resp, err := http.DefaultClient.Do(&req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("arc-raiders http %d: %s", resp.StatusCode, string(body))
	}

	var eventTimersResponse EventTimersResponse
	if err := json.Unmarshal(body, &eventTimersResponse); err != nil {
		return nil, fmt.Errorf("unmarshal event timers: %w", err)
	}

	return eventTimersResponse.Data, nil
}
