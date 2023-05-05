package google

import (
	"context"
	"time"
)

type Calendar interface {
	CreateMeet(ctx context.Context, req *CreateMeetReq) (*CreateMeetRsp, error)
	AddGuestsToMeet(ctx context.Context, req *AddGuestsToMeetReq) error
}

type AddGuestsToMeetReq struct {
	CalendarID string
	EventID    string
	Guests     []string
}

type CreateMeetReq struct {
	Summary     string
	CalendarID  string
	Description string
	Location    string
	Guests      []string
	StartTime   time.Time
	EndTime     time.Time
}

type CreateMeetRsp struct {
	EventID  string
	MeetLink string
}
