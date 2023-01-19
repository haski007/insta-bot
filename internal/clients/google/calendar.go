package google

import (
	"context"
	"time"
)

type Calendar interface {
	CreateMeet(ctx context.Context, req *CreateMeetReq) (*CreateMeetRsp, error)
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
	MeetLink string
}
