package listener

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"strings"

	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

const (
	anglicismCardWidth        = 720
	anglicismCardPadding      = 36
	anglicismCardAccentWidth  = 6
	anglicismCardAccentGap    = 14
	anglicismCardFontSize     = 28
	anglicismCardLineHeight   = 44
	anglicismCardMaxLines     = 22
	anglicismCardTextOffsetX  = anglicismCardPadding + anglicismCardAccentWidth + anglicismCardAccentGap
	anglicismCardTextMaxWidth = anglicismCardWidth - anglicismCardTextOffsetX - anglicismCardPadding
)

var (
	anglicismCardBg           = color.RGBA{R: 0x10, G: 0x16, B: 0x21, A: 0xff}
	anglicismCardFg           = color.RGBA{R: 0xf4, G: 0xf4, B: 0xf5, A: 0xff}
	anglicismCardAccentBlue   = color.RGBA{R: 0x00, G: 0x57, B: 0xb7, A: 0xff}
	anglicismCardAccentYellow = color.RGBA{R: 0xff, G: 0xd7, B: 0x00, A: 0xff}
)

var anglicismCardFace font.Face

func init() {
	parsed, err := opentype.Parse(goregular.TTF)
	if err != nil {
		panic("anglicism card: parse font: " + err.Error())
	}
	face, err := opentype.NewFace(parsed, &opentype.FaceOptions{
		Size:    anglicismCardFontSize,
		DPI:     96,
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic("anglicism card: face: " + err.Error())
	}
	anglicismCardFace = face
}

// renderAnglicismCard returns a PNG-encoded card with the correction text baked in.
func renderAnglicismCard(text string) ([]byte, error) {
	lines := wrapAnglicismCardText(text)
	if len(lines) == 0 {
		lines = []string{""}
	}
	if len(lines) > anglicismCardMaxLines {
		lines = lines[:anglicismCardMaxLines]
		lines[anglicismCardMaxLines-1] = truncateLineToFit(lines[anglicismCardMaxLines-1] + " …")
	}

	height := anglicismCardPadding*2 + anglicismCardLineHeight*len(lines)
	img := image.NewRGBA(image.Rect(0, 0, anglicismCardWidth, height))
	draw.Draw(img, img.Bounds(), &image.Uniform{anglicismCardBg}, image.Point{}, draw.Src)

	half := anglicismCardPadding + (height-anglicismCardPadding*2)/2
	accentLeft := anglicismCardPadding
	accentRight := accentLeft + anglicismCardAccentWidth
	draw.Draw(img,
		image.Rect(accentLeft, anglicismCardPadding, accentRight, half),
		&image.Uniform{anglicismCardAccentBlue}, image.Point{}, draw.Src)
	draw.Draw(img,
		image.Rect(accentLeft, half, accentRight, height-anglicismCardPadding),
		&image.Uniform{anglicismCardAccentYellow}, image.Point{}, draw.Src)

	drawer := &font.Drawer{
		Dst:  img,
		Src:  &image.Uniform{anglicismCardFg},
		Face: anglicismCardFace,
	}

	y := anglicismCardPadding + anglicismCardFontSize + 4
	for _, line := range lines {
		drawer.Dot = fixed.Point26_6{
			X: fixed.I(anglicismCardTextOffsetX),
			Y: fixed.I(y),
		}
		drawer.DrawString(line)
		y += anglicismCardLineHeight
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func wrapAnglicismCardText(text string) []string {
	var out []string
	for _, paragraph := range strings.Split(text, "\n") {
		if strings.TrimSpace(paragraph) == "" {
			out = append(out, "")
			continue
		}
		out = append(out, wrapAnglicismCardLine(paragraph)...)
	}
	return out
}

func wrapAnglicismCardLine(line string) []string {
	words := strings.Fields(line)
	if len(words) == 0 {
		return []string{""}
	}
	var lines []string
	cur := ""
	for _, w := range words {
		candidate := w
		if cur != "" {
			candidate = cur + " " + w
		}
		if font.MeasureString(anglicismCardFace, candidate).Round() <= anglicismCardTextMaxWidth {
			cur = candidate
			continue
		}
		if cur != "" {
			lines = append(lines, cur)
			cur = ""
		}
		if font.MeasureString(anglicismCardFace, w).Round() > anglicismCardTextMaxWidth {
			chunks := hardBreakByRunes(w)
			lines = append(lines, chunks[:len(chunks)-1]...)
			cur = chunks[len(chunks)-1]
		} else {
			cur = w
		}
	}
	if cur != "" {
		lines = append(lines, cur)
	}
	return lines
}

func hardBreakByRunes(s string) []string {
	var out []string
	cur := ""
	for _, r := range s {
		candidate := cur + string(r)
		if font.MeasureString(anglicismCardFace, candidate).Round() > anglicismCardTextMaxWidth && cur != "" {
			out = append(out, cur)
			cur = string(r)
			continue
		}
		cur = candidate
	}
	if cur != "" {
		out = append(out, cur)
	}
	if len(out) == 0 {
		out = []string{s}
	}
	return out
}

func truncateLineToFit(s string) string {
	if font.MeasureString(anglicismCardFace, s).Round() <= anglicismCardTextMaxWidth {
		return s
	}
	runes := []rune(s)
	for len(runes) > 1 {
		runes = runes[:len(runes)-1]
		candidate := string(runes) + "…"
		if font.MeasureString(anglicismCardFace, candidate).Round() <= anglicismCardTextMaxWidth {
			return candidate
		}
	}
	return string(runes)
}
