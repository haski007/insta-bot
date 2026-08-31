package listener

import (
	"testing"

	"github.com/haski007/insta-bot/internal/clients/instloader"
	"github.com/haski007/insta-bot/internal/clients/tiktokapi"
)

func TestExtractTikTokURL(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{
			name: "www video",
			text: "див https://www.tiktok.com/@user/video/1234567890",
			want: "https://www.tiktok.com/@user/video/1234567890",
		},
		{
			name: "share vm",
			text: "https://vm.tiktok.com/ZMxxxx/",
			want: "https://vm.tiktok.com/ZMxxxx/",
		},
		{
			name: "share vt",
			text: "look https://vt.tiktok.com/ZSxxxx/ please",
			want: "https://vt.tiktok.com/ZSxxxx/",
		},
		{
			name: "photo post",
			text: "https://www.tiktok.com/@user/photo/7375880582473043232",
			want: "https://www.tiktok.com/@user/photo/7375880582473043232",
		},
		{
			name: "mobile",
			text: "https://m.tiktok.com/v/1234567890.html",
			want: "https://m.tiktok.com/v/1234567890.html",
		},
		{
			name: "trailing punctuation",
			text: "watch https://vm.tiktok.com/ZMxxxx/.",
			want: "https://vm.tiktok.com/ZMxxxx/",
		},
		{
			name: "not tiktok",
			text: "https://www.instagram.com/reel/abc/",
			want: "",
		},
		{
			name: "fake host",
			text: "https://faketiktok.com/video/1",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractTikTokURL(tt.text)
			if got != tt.want {
				t.Errorf("extractTikTokURL() = %q, want %q", got, tt.want)
			}
			if (got != "") != isTikTokURL(tt.text) {
				t.Errorf("isTikTokURL() mismatch for %q", tt.text)
			}
		})
	}
}

func TestTiktokMediaItems(t *testing.T) {
	video := tiktokMediaItems(&tiktokapi.Media{VideoURL: "https://cdn.example/v.mp4"})
	if len(video) != 1 || !video[0].IsVideo || video[0].VideoURL != "https://cdn.example/v.mp4" {
		t.Fatalf("video items = %+v", video)
	}

	photos := tiktokMediaItems(&tiktokapi.Media{
		Images: []string{"https://cdn.example/1.jpg", "https://cdn.example/2.jpg"},
	})
	if len(photos) != 2 {
		t.Fatalf("photo items = %+v", photos)
	}
	if photos[0] != (instloader.MediaItem{URL: "https://cdn.example/1.jpg"}) {
		t.Errorf("first photo = %+v", photos[0])
	}

	if items := tiktokMediaItems(&tiktokapi.Media{}); len(items) != 0 {
		t.Errorf("empty media should be empty, got %+v", items)
	}
}
