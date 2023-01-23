package listener

import (
	"context"
	"fmt"

	"github.com/haski007/insta-bot/internal/clients/google"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func (rcv *InstaBotService) triggerPollAnswer(update tgbotapi.Update) {
	username := update.PollAnswer.User.UserName
	pollID := update.PollAnswer.PollID

	if update.PollAnswer.OptionIDs == nil && len(update.PollAnswer.OptionIDs) < 1 {
		return
	}

	found, err := rcv.storage.PollExists(pollID)
	if err != nil {
		rcv.NotifyCreator(fmt.Sprintf("[triggerPollAnswer] poll exists err: %s", err))
		logrus.WithError(err).Println("[triggerPollAnswer] poll exists")
		return
	}

	userEmail, err := rcv.storage.GetUser(username)
	if err != nil {
		rcv.NotifyCreator(fmt.Sprintf("[triggerPollAnswer] getUser err: %s", err))
		logrus.WithError(err).WithField("username", username).Println("[triggerPollAnswer] getUser")
		return
	}

	if found && update.PollAnswer.OptionIDs[0] == 0 && userEmail != "" {
		guests := []string{userEmail}

		poll, err := rcv.storage.GetPoll(pollID)
		if err != nil {
			rcv.NotifyCreator(fmt.Sprintf("[triggerPollAnswer] get poll err: %s", err))
			logrus.WithError(err).Println("[triggerPollAnswer] get poll")
			return
		}

		if err := rcv.calendar.AddGuestsToMeet(context.Background(), &google.AddGuestsToMeetReq{
			CalendarID: "primary",
			EventID:    poll.GoogleEventID,
			Guests:     guests,
		}); err != nil {
			rcv.NotifyCreator(fmt.Sprintf("[triggerPollAnswer] add guests to a meet err: %s", err))
			logrus.WithError(err).Println("[triggerPollAnswer] add guests to a meet")
			return
		}

		rcv.log.WithFields(map[string]interface{}{
			"username": username,
			"pollID":   pollID,
			"guests":   guests,
		}).
			Debugln("added voted yes to an event guests")
	}
}
