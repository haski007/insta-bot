package instapi

import (
	"strings"
)

type Api struct {
}

func New() *Api {
	return &Api{}
}

func getSubstrBefore(value string, a string) string {
	// Get substring before a string.
	pos := strings.Index(value, a)
	if pos == -1 {
		return ""
	}
	return value[0:pos]
}
