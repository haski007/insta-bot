package listener

import (
	"strings"
	"time"
)

func (rcv *InstaBotService) RedisMonitor() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		readonly, err := rcv.storage.IsReadOnly()
		if err != nil {
			rcv.NotifyCreator("[RedisMonitor] " + err.Error())
			continue
		}

		if readonly {
			ticker.Reset(6 * time.Hour)
			rcv.log.Debugln("[RedisMonitor] redis is readonly")
		} else {
			rcv.log.Debugln("[RedisMonitor] redis is NOT readonly")
		}
	}
}

func escapeMarkdown(v string) string {
	v = strings.ReplaceAll(v, "_", "\\_")
	v = strings.ReplaceAll(v, "*", "\\*")
	v = strings.ReplaceAll(v, "[", "\\[")
	v = strings.ReplaceAll(v, "]", "\\]")
	v = strings.ReplaceAll(v, "~", "\\~")
	v = strings.ReplaceAll(v, "`", "\\`")
	v = strings.ReplaceAll(v, ">", "\\>")
	v = strings.ReplaceAll(v, "+", "\\+")
	v = strings.ReplaceAll(v, "-", "\\-")
	v = strings.ReplaceAll(v, "=", "\\=")
	v = strings.ReplaceAll(v, "{", "\\{")
	v = strings.ReplaceAll(v, "}", "\\}")
	v = strings.ReplaceAll(v, "|", "\\|")
	v = strings.ReplaceAll(v, ".", "\\.")
	v = strings.ReplaceAll(v, "!", "\\!")

	return v
}
