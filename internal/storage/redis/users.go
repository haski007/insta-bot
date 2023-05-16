package redis

import (
	"errors"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/haski007/insta-bot/internal/storage"
)

func (r *redisClient) AddUser(username, email string) error {
	return r.conn.Set(r.keyFromUser(username), email, 0).Err()
}

func (r *redisClient) GetUser(username string) (email string, err error) {
	email, err = r.conn.Get(r.keyFromUser(username)).Result()
	if err != nil && errors.Is(err, redis.Nil) {
		return "", storage.ErrNotFound
	}

	return
}

func (r *redisClient) keyFromUser(username string) string {
	return fmt.Sprintf("users/%s", username)
}
