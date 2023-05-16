package redis

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/haski007/insta-bot/internal/storage"
)

type redisClient struct {
	conn    *redis.Client
	convTTL time.Duration
}

func NewClient(conn *redis.Client, convTTL time.Duration) (storage.Storage, error) {
	if err := conn.Ping().Err(); err != nil {
		return nil, err
	}
	return &redisClient{
		conn:    conn,
		convTTL: convTTL,
	}, nil
}

func (r *redisClient) IsReadOnly() (readOnly bool, err error) {
	if err := r.conn.Set("test-write-rights", true, time.Second).Err(); err != nil {
		return true, err
	}

	return false, nil
}
