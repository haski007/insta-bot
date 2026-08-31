package tiktokapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMedia_video(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("url") == "" {
			t.Errorf("expected url query param")
		}
		if r.URL.Query().Get("hd") != "1" {
			t.Errorf("expected hd=1, got %q", r.URL.Query().Get("hd"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"code": 0,
			"msg": "success",
			"data": {
				"id": "7106658991907802411",
				"title": "which biome",
				"play": "https://cdn.example/play.mp4",
				"hdplay": "/video/media/hd.mp4",
				"hd_size": 4000000,
				"size": 2000000,
				"cover": "https://cdn.example/cover.jpg",
				"author": {"unique_id": "tiktok", "nickname": "TikTok"}
			}
		}`))
	}))
	defer server.Close()

	client := &TikTokClient{apiBase: server.URL, http: server.Client()}
	media, err := client.GetMedia("https://www.tiktok.com/@tiktok/video/7106658991907802411")
	if err != nil {
		t.Fatalf("GetMedia: %v", err)
	}
	if media.ID != "7106658991907802411" {
		t.Errorf("id = %q", media.ID)
	}
	if media.Author != "tiktok" {
		t.Errorf("author = %q", media.Author)
	}
	if media.Title != "which biome" {
		t.Errorf("title = %q", media.Title)
	}
	if media.VideoURL != "https://www.tikwm.com/video/media/hd.mp4" {
		t.Errorf("video url = %q", media.VideoURL)
	}
	if len(media.Images) != 0 {
		t.Errorf("expected no images, got %v", media.Images)
	}
}

func TestGetMedia_photos(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"code": 0,
			"data": {
				"id": "7375880582473043232",
				"title": "slides",
				"images": ["/photo/1.jpg", "https://cdn.example/2.jpg"],
				"author": {"unique_id": "someone"}
			}
		}`))
	}))
	defer server.Close()

	client := &TikTokClient{apiBase: server.URL, http: server.Client()}
	media, err := client.GetMedia("https://www.tiktok.com/@someone/photo/7375880582473043232")
	if err != nil {
		t.Fatalf("GetMedia: %v", err)
	}
	if len(media.Images) != 2 {
		t.Fatalf("images = %v", media.Images)
	}
	if media.Images[0] != "https://www.tikwm.com/photo/1.jpg" {
		t.Errorf("image 0 = %q", media.Images[0])
	}
	if media.Images[1] != "https://cdn.example/2.jpg" {
		t.Errorf("image 1 = %q", media.Images[1])
	}
}

func TestGetMedia_apiError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code": -1, "msg": "Url parsing is failed!"}`))
	}))
	defer server.Close()

	client := &TikTokClient{apiBase: server.URL, http: server.Client()}
	_, err := client.GetMedia("https://www.tiktok.com/@x/video/1")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestPickPlayURL_skipsHugeHD(t *testing.T) {
	got := pickPlayURL(&apiData{
		HDPlay: "https://cdn.example/hd.mp4",
		Play:   "https://cdn.example/play.mp4",
		HDSize: telegramBotFileLimit + 1,
		Size:   2_000_000,
	})
	if got != "https://cdn.example/play.mp4" {
		t.Errorf("got %q", got)
	}
}
