package calendar

import (
	"context"
	"time"

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
	var attendees []*calendar.EventAttendee

	for _, email := range req.Guests {
		attendees = append(attendees, &calendar.EventAttendee{Email: email})
	}

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

		Attendees: attendees,
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
	}, nil
}
