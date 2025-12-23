package listener

import (
	"fmt"
	"strings"
	"time"

	arcraiders "github.com/haski007/insta-bot/internal/clients/arc-raiders"
)

const (
	MatriarchEventName = "Matriarch"
	HarvesterEventName = "Harvester"
	NightRaidEventName = "Night Raid"

	StellaMontisMapName = "Stella Montis"

	ARCMonitorTime = "17:00"
)

// everyday at defined time  check if there are any ARC events happening today
func (rcv *InstaBotService) RunARCMonitor() {
	location, err := time.LoadLocation("Europe/Warsaw")
	if err != nil {
		panic(err)
	}

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	timeToCheck, err := time.Parse("15:04", ARCMonitorTime)
	if err != nil {
		panic(err)
	}

LOOP:
	for {
		select {
		case <-rcv.ctx.Done():
			rcv.log.Info("ARC Monitor stopped")
			break LOOP
		case t := <-ticker.C:
			if t.In(location).Hour() == timeToCheck.Hour() && t.In(location).Minute() == timeToCheck.Minute() {
				if err := rcv.checkAndNotifyARCEvents(); err != nil {
					rcv.log.WithError(err).Error("[RunARCMonitor] checkAndNotifyARCEvents")
					rcv.NotifyCreator(fmt.Sprintf("[RunARCMonitor] checkAndNotifyARCEvents err: %s\n", err))
				}
			}
		}
	}
}

func (rcv *InstaBotService) checkAndNotifyARCEvents() error {
	// Get all events
	events, err := rcv.arcRaidersClient.GetARCEvents(nil)
	if err != nil {
		return fmt.Errorf("get ARC events: %w", err)
	}

	// Filter events: Matriarch, Harvester, or Night Raid on Stella Montis
	var filteredEvents []arcraiders.EventTimer
	for _, event := range events {
		// Matriarch (any map)
		if strings.EqualFold(event.Name, MatriarchEventName) {
			filteredEvents = append(filteredEvents, event)
			continue
		}
		// Harvester (any map)
		if strings.EqualFold(event.Name, HarvesterEventName) {
			filteredEvents = append(filteredEvents, event)
			continue
		}
		// Night Raid on Stella Montis only
		if strings.EqualFold(event.Name, NightRaidEventName) && strings.EqualFold(event.Map, StellaMontisMapName) {
			filteredEvents = append(filteredEvents, event)
			continue
		}
	}

	// If no matching events, return early
	if len(filteredEvents) == 0 {
		rcv.log.Info("[checkAndNotifyARCEvents] no matching events (Matriarch/Harvester/Night Raid on Stella Montis)")
		return nil
	}

	// Get all subscribed chats
	chatIDs, err := rcv.storage.GetAllARCSubscribedChats()
	if err != nil {
		return fmt.Errorf("get all ARC subscribed chats: %w", err)
	}

	if len(chatIDs) == 0 {
		rcv.log.Info("[checkAndNotifyARCEvents] no subscribed chats")
		return nil
	}

	// Build message
	var messageBuilder strings.Builder
	messageBuilder.WriteString("🎮 *ARC Raiders Events Today*\n\n")

	for _, event := range filteredEvents {
		messageBuilder.WriteString(fmt.Sprintf("*%s*\n", event.Name))
		if event.Map != "" {
			messageBuilder.WriteString(fmt.Sprintf("📍 Map: %s\n", event.Map))
		}
		if event.Description != "" {
			messageBuilder.WriteString(fmt.Sprintf("📝 %s\n", event.Description))
		}
		if len(event.Times) > 0 {
			var timeStrings []string
			for _, tr := range event.Times {
				timeStrings = append(timeStrings, fmt.Sprintf("%s-%s", tr.Start, tr.End))
			}
			messageBuilder.WriteString(fmt.Sprintf("⏰ Times: %s\n", strings.Join(timeStrings, ", ")))
		}
		messageBuilder.WriteString("\n")
	}

	message := messageBuilder.String()

	// Send to all subscribed chats
	for _, chatID := range chatIDs {
		if err := rcv.SendMessage(chatID, message); err != nil {
			rcv.log.WithError(err).WithField("chat_id", chatID).Error("[checkAndNotifyARCEvents] send message")
			// Continue with other chats even if one fails
		}
	}

	rcv.log.WithField("events_count", len(filteredEvents)).
		WithField("chats_notified", len(chatIDs)).
		Info("[checkAndNotifyARCEvents] notifications sent")

	return nil
}
