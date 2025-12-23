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

	ARCMonitorTime = "18:00"
)

// everyday at defined time check if there are any ARC events happening today (Kyiv time)
func (rcv *InstaBotService) RunARCMonitor() {
	location, err := time.LoadLocation("Europe/Kiev")
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
	filteredEvents := rcv.filterARCEvents(events)

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

	message, err := rcv.buildARCEventMessage(filteredEvents)
	if err != nil {
		return fmt.Errorf("build ARC event message: %w", err)
	}

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

func (rcv *InstaBotService) filterARCEvents(events []arcraiders.EventTimer) []arcraiders.EventTimer {
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
	return filteredEvents
}

func (rcv *InstaBotService) buildARCEventMessage(events []arcraiders.EventTimer) (string, error) {
	// Build message
	var messageBuilder strings.Builder
	messageBuilder.WriteString("🎮 *ARC Raiders Events Today (Kyiv time)*\n\n")

	// Prepare Kyiv location and today's date (used to translate from UTC-like schedule to Kyiv time)
	kyivLocation, err := time.LoadLocation("Europe/Kiev")
	if err != nil {
		return "", fmt.Errorf("load Kyiv location: %w", err)
	}
	nowUTC := time.Now().UTC()

	for _, event := range events {
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
				startParsed, errStart := time.Parse("15:04", tr.Start)
				endParsed, errEnd := time.Parse("15:04", tr.End)
				if errStart != nil || errEnd != nil {
					// fallback to raw values if parsing fails
					timeStrings = append(timeStrings, fmt.Sprintf("%s-%s", tr.Start, tr.End))
					continue
				}

				startUTC := time.Date(
					nowUTC.Year(), nowUTC.Month(), nowUTC.Day(),
					startParsed.Hour(), startParsed.Minute(), 0, 0, time.UTC,
				)
				endUTC := time.Date(
					nowUTC.Year(), nowUTC.Month(), nowUTC.Day(),
					endParsed.Hour(), endParsed.Minute(), 0, 0, time.UTC,
				)

				startKyiv := startUTC.In(kyivLocation)
				endKyiv := endUTC.In(kyivLocation)

				timeStrings = append(timeStrings, fmt.Sprintf("%s-%s", startKyiv.Format("15:04"), endKyiv.Format("15:04")))
			}
			messageBuilder.WriteString(fmt.Sprintf("⏰ Times (Kyiv): %s\n", strings.Join(timeStrings, ", ")))
		}
		messageBuilder.WriteString("\n")
	}

	return messageBuilder.String(), nil
}
