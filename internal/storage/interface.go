package storage

type Storage interface {
	AddChatWithMembers(context string, chatID int64, members []string) error
	GetChatMembers(context string, chatID int64) (members []string, err error)

	ChatExists(context string, chatID int64) (bool, error)
	DeleteChat(context string, chatID int64) error

	GetAllCTXs(chatID int64) (map[string][]string, error)
}
