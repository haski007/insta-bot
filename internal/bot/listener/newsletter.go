package listener

import (
	"fmt"
	"time"
)

func (rcv *InstaBotService) RunNewsLetter() {
	location, err := time.LoadLocation("Local") // Replace "Local" with your timezone if needed
	if err != nil {
		panic(err)
	}

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

LOOP:
	for {
		select {
		case <-rcv.ctx.Done():
			rcv.log.Info("NewsLetter stopped")
			break LOOP
		case t := <-ticker.C:
			if t.In(location).Hour() == 19 && t.Minute() == 00 {
				if err := rcv.sendStartupNewsletter(); err != nil {
					rcv.log.WithError(err).Error("[RunNewsLetter] sendStartupNewsletter")
					rcv.NotifyCreator(fmt.Sprintf("[RunNewsLetter] sendStartupNewsletter err: %s\n", err))
				}
			}
		}
	}
}