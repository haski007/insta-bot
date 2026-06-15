package listener

import (
	"strings"
	"time"
)

func (rcv *InstaBotService) RedisMonitor() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	var lastReadonly *bool

	for range ticker.C {
		readonly, err := rcv.storage.IsReadOnly()
		if err != nil {
			rcv.NotifyCreator("[RedisMonitor] " + err.Error())
			continue
		}

		if lastReadonly != nil && *lastReadonly == readonly {
			continue
		}

		wasReadonly := lastReadonly != nil && *lastReadonly
		v := readonly
		lastReadonly = &v

		if readonly {
			ticker.Reset(6 * time.Hour)
			rcv.log.Warnln("[RedisMonitor] redis is readonly")
			continue
		}

		ticker.Reset(time.Minute)
		if wasReadonly {
			rcv.log.Infoln("[RedisMonitor] redis is writable again")
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
