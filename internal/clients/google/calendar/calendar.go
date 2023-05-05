package calendar

import (
	"context"
	"fmt"
	"time"

	"github.com/haski007/insta-bot/internal/clients/google/transform"

	"github.com/google/uuid"
	"github.com/haski007/insta-bot/internal/clients/google"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/calendar/v3"
)

type wrapper struct {
	srv *calendar.Service
}

func New(srv *calendar.Service) google.Calendar {
	return &wrapper{srv: srv}
}

//Summary:     "Test cs go meeting",
//Location:    "Forsage dance club",
//Description: "Here we gonna make some magic and win all the enemies!",

func (rcv *wrapper) CreateMeet(ctx context.Context, req *google.CreateMeetReq) (*google.CreateMeetRsp, error) {
	event := &calendar.Event{
		Summary:     req.Summary,
		Location:    req.Location,
		Description: req.Description,
		Start: &calendar.EventDateTime{
			DateTime: req.StartTime.Format(time.RFC3339),
			TimeZone: "UTC",
		},
		End: &calendar.EventDateTime{
			DateTime: req.EndTime.Format(time.RFC3339),
			TimeZone: "UTC",
		},
		ConferenceData: &calendar.ConferenceData{
			CreateRequest: &calendar.CreateConferenceRequest{
				RequestId: uuid.New().String(),
			},
		},

		Attendees: transform.ToAttendees(req.Guests),
	}

	// Insert the Meet into the calendar
	calendarId := req.CalendarID
	createdEvent, err := rcv.srv.Events.Insert(calendarId, event).
		ConferenceDataVersion(1).
		Context(ctx).
		Do()
	if err != nil {
		logrus.WithError(err).Fatal("insert event")
	}

	return &google.CreateMeetRsp{
		MeetLink: createdEvent.HangoutLink,
		EventID:  createdEvent.Id,
	}, nil
}

func (rcv *wrapper) AddGuestsToMeet(_ context.Context, req *google.AddGuestsToMeetReq) error {
	event, err := rcv.srv.Events.Get(req.CalendarID, req.EventID).Do()
	if err != nil {
		return fmt.Errorf("get event err: %w", err)
	}

	event.Attendees = append(event.Attendees, transform.ToAttendees(req.Guests)...)

	_, err = rcv.srv.Events.Update(req.CalendarID, req.EventID, event).Do()
	return err
}
