package listener

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// renderAnglicismCardVideo renders the text card as PNG and wraps it into a short MP4.
// Requires ffmpeg in PATH (installed in the container image).
func renderAnglicismCardVideo(text string) ([]byte, error) {
	pngBytes, err := renderAnglicismCard(text)
	if err != nil {
		return nil, fmt.Errorf("render card: %w", err)
	}
	return wrapPNGAsMP4(pngBytes)
}

func wrapPNGAsMP4(pngBytes []byte) ([]byte, error) {
	dir, err := os.MkdirTemp("", "anglicism-vid-")
	if err != nil {
		return nil, fmt.Errorf("mktemp dir: %w", err)
	}
	defer os.RemoveAll(dir)

	inPath := filepath.Join(dir, "card.png")
	outPath := filepath.Join(dir, "card.mp4")
	if err := os.WriteFile(inPath, pngBytes, 0o600); err != nil {
		return nil, fmt.Errorf("write png: %w", err)
	}

	var stderr bytes.Buffer
	cmd := exec.Command("ffmpeg",
		"-y",
		"-loop", "1",
		"-i", inPath,
		"-t", "2",
		"-c:v", "libx264",
		"-preset", "ultrafast",
		"-tune", "stillimage",
		"-pix_fmt", "yuv420p",
		"-movflags", "+faststart",
		"-an",
		outPath,
	)
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("ffmpeg: %w (%s)", err, stderr.String())
	}

	return os.ReadFile(outPath)
}
