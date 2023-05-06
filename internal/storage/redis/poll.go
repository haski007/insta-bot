package redis

import (
	"fmt"
	"strings"
	"time"

	"github.com/haski007/insta-bot/internal/storage"
)

func (r *redisClient) keyFromPollID(pollID string) string {
	return fmt.Sprintf("poll/%s", pollID)
}

func (r *redisClient) PollExists(pollID string) (bool, error) {
	exists, err := r.conn.Exists(r.keyFromPollID(pollID)).Result()
	if err != nil {
		return false, fmt.Errorf("redis get bool err: %w", err)
	}
	return exists == 1, nil
}

func (r *redisClient) PollInit(poll storage.Poll) error {
	return r.conn.Set(r.keyFromPollID(poll.ID), poll, time.Hour*24).Err()
}

func (r *redisClient) GetPoll(pollID string) (*storage.Poll, error) {
	var poll = new(storage.Poll)

	if err := r.conn.Get(r.keyFromPollID(pollID)).Scan(poll); err != nil {
		return nil, fmt.Errorf("get and scan poll key: %s err: %w", r.keyFromPollID(pollID), err)
	}

	return poll, nil
}

func (r *redisClient) GetAllPolls() (map[string]*storage.Poll, error) {
	var result = make(map[string]*storage.Poll)
	iter := r.conn.Scan(0, fmt.Sprintf("poll/*"), 0).Iterator()
	for iter.Next() {
		poll, err := r.GetPoll(strings.Split(iter.Val(), "/")[1])
		if err != nil {
			return nil, fmt.Errorf("get pull err: %w", err)
		}

		result[iter.Val()] = poll
	}

	return result, nil
}
