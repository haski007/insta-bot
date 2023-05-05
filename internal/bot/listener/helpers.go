package listener

import (
	"time"
)

func (rcv *InstaBotService) RedisMonitor() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		readonly, err := rcv.storage.IsReadOnly()
		if err != nil {
			rcv.NotifyCreator("[RedisMonitor] " + err.Error())
			continue
		}

		if readonly {
			ticker.Reset(6 * time.Hour)
			rcv.NotifyCreator("[RedisMonitor] Redis is in read-only mode")
		} else {
			rcv.log.WithField("readonly", readonly).Debugln("[RedisMonitor] redis is NOT readonly")
		}
	}

}
