package text

import "strings"

// DecodeUrl - decodes characters:
// \u002F => '/',  %7C => '|',  %3D => '='
func DecodeUrl(url string) (decoded string) {
	decoded = strings.ReplaceAll(url, "\\u002F", "/")
	decoded = strings.ReplaceAll(url, "\\u0026", "&")
	decoded = strings.ReplaceAll(decoded, "%7C", "|")
	decoded = strings.ReplaceAll(decoded, "%3D", "=")

	return decoded
}
