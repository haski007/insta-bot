package youtube

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestGetMedia_success(t *testing.T) {
	var progressHits int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case strings.Contains(r.URL.Path, "download"):
			if r.URL.Query().Get("url") == "" {
				t.Errorf("expected url query param")
			}
			if r.URL.Query().Get("format") != "720" {
				t.Errorf("expected format=720, got %q", r.URL.Query().Get("format"))
			}
			_ = json.NewEncoder(w).Encode(map[string]any{
				"id":            "job-1",
				"title":         "Demo Short",
				"success":       true,
				"progress_url":  "http://" + r.Host + "/progress",
				"thumbnail_url": "https://cdn.example/thumb.jpg",
			})
		case strings.Contains(r.URL.Path, "progress"):
			progressHits++
			if progressHits < 2 {
				_ = json.NewEncoder(w).Encode(map[string]any{
					"progress": 50,
					"success":  0,
					"text":     "Preparing",
				})
				return
			}
			_ = json.NewEncoder(w).Encode(map[string]any{
				"progress":      1000,
				"success":       1,
				"title":         "Demo Short",
				"download_url":  "https://cdn.example/video.mp4",
				"thumbnail_url": "https://cdn.example/thumb.jpg",
			})
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	client := &Client{
		downloadBase: server.URL + "/download",
		http:         server.Client(),
		maxQuality:   720,
		pollInterval: 10 * time.Millisecond,
		pollTimeout:  time.Second,
	}

	media, err := client.GetMedia("https://www.youtube.com/shorts/abc123")
	if err != nil {
		t.Fatalf("GetMedia: %v", err)
	}
	if media.Title != "Demo Short" {
		t.Errorf("title = %q", media.Title)
	}
	if media.VideoURL != "https://cdn.example/video.mp4" {
		t.Errorf("video url = %q", media.VideoURL)
	}
	if media.ID != "job-1" {
		t.Errorf("id = %q", media.ID)
	}
	if media.Quality != "720" {
		t.Errorf("quality = %q", media.Quality)
	}
	if progressHits < 2 {
		t.Errorf("expected polling, hits=%d", progressHits)
	}
}

func TestGetMedia_noDownloadURL(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if strings.Contains(r.URL.Path, "download") {
			_ = json.NewEncoder(w).Encode(map[string]any{
				"success":      true,
				"progress_url": "http://" + r.Host + "/progress",
			})
			return
		}
		_ = json.NewEncoder(w).Encode(map[string]any{
			"progress": 1000,
			"success":  1,
			"text":     "Finished",
		})
	}))
	defer server.Close()

	client := &Client{
		downloadBase: server.URL + "/download",
		http:         server.Client(),
		maxQuality:   1080,
		pollInterval: 5 * time.Millisecond,
		pollTimeout:  200 * time.Millisecond,
	}
	_, err := client.GetMedia("https://www.youtube.com/shorts/missing")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestFormatQuality(t *testing.T) {
	cases := map[int]string{240: "360", 360: "360", 480: "480", 700: "720", 720: "720", 1080: "1080", 2160: "1080"}
	for in, want := range cases {
		if got := formatQuality(in); got != want {
			t.Errorf("formatQuality(%d)=%q want %q", in, got, want)
		}
	}
}
