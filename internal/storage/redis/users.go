package redis

import (
	"fmt"

	"github.com/haski007/insta-bot/internal/bot/listener"
)

func (r *redisClient) AddUser(username, email string) error {
	return r.conn.Set(r.keyFromUser(username), email, 0).Err()
}

func (r *redisClient) GetUser(username string) (email string, err error) {
	return r.conn.Get(r.keyFromUser(username)).Result()
}

func (r *redisClient) keyFromUser(username string) string {
	return fmt.Sprintf("%s/%s", listener.UsersContext, username)
}
