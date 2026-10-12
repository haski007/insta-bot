package listener

import "testing"

func TestExtractYouTubeShortsURL(t *testing.T) {
	cases := []struct {
		name string
		text string
		want string
	}{
		{
			name: "www shorts",
			text: "див https://www.youtube.com/shorts/abc123XYZ",
			want: "https://www.youtube.com/shorts/abc123XYZ",
		},
		{
			name: "bare youtube shorts",
			text: "https://youtube.com/shorts/abc123XYZ",
			want: "https://youtube.com/shorts/abc123XYZ",
		},
		{
			name: "m.youtube shorts with punctuation",
			text: "look https://m.youtube.com/shorts/abc123XYZ/ please",
			want: "https://m.youtube.com/shorts/abc123XYZ/",
		},
		{
			name: "trailing punctuation",
			text: "https://www.youtube.com/shorts/abc123XYZ.",
			want: "https://www.youtube.com/shorts/abc123XYZ",
		},
		{
			name: "watch url ignored",
			text: "https://www.youtube.com/watch?v=abc123XYZ",
			want: "",
		},
		{
			name: "youtu.be ignored",
			text: "https://youtu.be/abc123XYZ",
			want: "",
		},
		{
			name: "not youtube",
			text: "https://example.com/shorts/abc",
			want: "",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := extractYouTubeShortsURL(tc.text); got != tc.want {
				t.Errorf("extractYouTubeShortsURL(%q)=%q want %q", tc.text, got, tc.want)
			}
		})
	}
}
