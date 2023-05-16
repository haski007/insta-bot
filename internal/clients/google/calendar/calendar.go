package calendar

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/haski007/insta-bot/internal/clients/google"
	"github.com/haski007/insta-bot/internal/clients/google/transform"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type wrapper struct {
	srv         *calendar.Service
	tokenSource oauth2.TokenSource
	config      *oauth2.Config
}

func New(
	srv *calendar.Service,
	tokenSource oauth2.TokenSource,
	config *oauth2.Config,
) google.Calendar {
	return &wrapper{
		srv:         srv,
		tokenSource: tokenSource,
		config:      config,
	}
}

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
		return nil, err
	}

	return &google.CreateMeetRsp{
		MeetLink: createdEvent.HangoutLink,
		EventID:  createdEvent.Id,
	}, nil
}

func (rcv *wrapper) RefreshToken(ctx context.Context) error {
	token, err := rcv.tokenSource.Token()
	if err != nil {
		return fmt.Errorf("token source err: %w", err)
	}
	rcv.srv, err = calendar.NewService(
		context.Background(),
		option.WithHTTPClient(rcv.config.Client(ctx, token)),
	)
	return nil
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
