package text

import (
	"strings"
	"unicode"
)

func CharLimiterToWord(src string, limit int) (limited string) {
	var lastWhitespaceIdx int
	runes := []rune(src)
	for idx, r := range runes {
		if unicode.IsSpace(r) {
			lastWhitespaceIdx = idx
		}

		if idx >= limit {
			if lastWhitespaceIdx == 0 {
				return string(runes[:idx]) + "..."
			}

			return strings.TrimSpace(string(runes[:lastWhitespaceIdx]) + "...")
		}
	}

	return src
}
