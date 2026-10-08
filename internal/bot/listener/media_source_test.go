package listener

import "testing"

func TestMediaDownloadSource(t *testing.T) {
	cases := map[string]string{
		"https://www.tikwm.com/video/foo.mp4":                         "tiktok",
		"https://v16-webapp-prime.tiktok.com/video/1":                 "tiktok",
		"https://scontent.cdninstagram.com/v/t51.2885-15/x.jpg":        "instagram",
		"https://instagram.fhel3-1.fna.fbcdn.net/v/t51.2885-15/x.jpg": "instagram",
		"https://example.com/file.mp4":                                "other",
		"://bad": "other",
	}
	for raw, want := range cases {
		if got := mediaDownloadSource(raw); got != want {
			t.Errorf("mediaDownloadSource(%q)=%q want %q", raw, got, want)
		}
	}
}
