package storage

import (
	"encoding/json"
	"time"
)

type Storage interface {
	AddChatWithMembers(context string, chatID int64, members []string) error
	GetChatMembers(context string, chatID int64) (members []string, err error)

	ChatExists(context string, chatID int64) (bool, error)
	DeleteChat(context string, chatID int64) error

	GetAllCTXs(chatID int64) (map[string][]string, error)

	PollInit(poll Poll) error
	PollExists(pollID string) (bool, error)
	GetPoll(pollID string) (*Poll, error)
	GetAllPolls() (map[string]*Poll, error)

	AddUser(username, email string) error
	GetUser(username string) (email string, err error)
}

type Poll struct {
	ID            string    `json:"id"`
	ChatID        int64     `json:"chat_id"`
	MeetLink      string    `json:"meet_link"`
	Time          time.Time `json:"time"`
	GoogleEventID string    `json:"google_event_id"`
}

func (p *Poll) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &p)
}

func (p Poll) MarshalBinary() (data []byte, err error) {
	return json.Marshal(p)
}
