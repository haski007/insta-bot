package storage

import (
	"encoding/json"
	"errors"
	"time"
)

var (
	ErrNotFound = errors.New("key not found")
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

	// Chat GPT
	SetSystemRoleForChat(chatID int64, role string) error
	GetSystemRole(chatID int64) (role string, err error)
	PushConversation(req *PushConversationReq) (err error)
	GetConversation(req *GetConversationReq) (conversation []Replica, err error)
	DropConversation(req *DropConversationReq) (err error)

	SubscribeChatToStartup(chatID int64) error
	UnsubscribeChatToStartup(chatID int64) error
	PushStartupNewsletter(chatID int64, replicas []Replica) error
	GetStartupNewsletter(chatID int64) (conversation []Replica, err error)
	GetAllStartupNewsletters() (newsletterConv map[int64][]Replica, err error)

	IsReadOnly() (bool, error)
}

type Replica struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	Name    string `json:"name"`
}

func (r *Replica) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &r)
}

func (r Replica) MarshalBinary() (data []byte, err error) {
	return json.Marshal(r)
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

type PushConversationReq struct {
	Username string    `json:"username"`
	UserID   int64     `json:"user_id"`
	ChatID   int64     `json:"chat_id"`
	Replicas []Replica `json:"replicas"`
}

type GetConversationReq struct {
	Username string `json:"username"`
	UserID   int64  `json:"user_id"`
	ChatID   int64  `json:"chat_id"`
}

type DropConversationReq struct {
	Username string `json:"username"`
	UserID   int64  `json:"user_id"`
	ChatID   int64  `json:"chat_id"`
}
