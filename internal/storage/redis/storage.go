package redis

import (
	"github.com/go-redis/redis"
	"github.com/haski007/insta-bot/internal/storage"
)

type redisClient struct {
	conn *redis.Client
}

func NewClient(conn *redis.Client) (storage.Storage, error) {
	if err := conn.Ping().Err(); err != nil {
		return nil, err
	}
	return &redisClient{
		conn: conn,
	}, nil
}
