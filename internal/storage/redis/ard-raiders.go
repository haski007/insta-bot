package redis

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	arcSubscribedKey = "arc_subscribed"
)

func (r *redisClient) SubscribeChatToARC(chatID int64) error {
	return r.conn.SAdd(r.keyFromChatID(arcSubscribedKey, chatID), chatID).Err()
}

func (r *redisClient) UnsubscribeChatToARC(chatID int64) error {
	return r.conn.SRem(r.keyFromChatID(arcSubscribedKey, chatID), chatID).Err()
}

func (r *redisClient) IsChatARCSubscribed(chatID int64) (bool, error) {
	exists, err := r.conn.SIsMember(r.keyFromChatID(arcSubscribedKey, chatID), chatID).Result()
	if err != nil {
		return false, fmt.Errorf("redis sismember err: %w", err)
	}
	return exists, nil
}

func (r *redisClient) GetAllARCSubscribedChats() ([]int64, error) {
	var chatIDs []int64
	pattern := fmt.Sprintf("*/%s", arcSubscribedKey)
	iter := r.conn.Scan(0, pattern, 0).Iterator()

	for iter.Next() {
		key := iter.Val()
		parts := strings.Split(key, "/")
		if len(parts) < 2 {
			continue
		}

		chatID, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			continue
		}

		// Check if chatID is actually in the set
		exists, err := r.conn.SIsMember(key, chatID).Result()
		if err != nil || !exists {
			continue
		}

		chatIDs = append(chatIDs, chatID)
	}

	if err := iter.Err(); err != nil {
		return nil, fmt.Errorf("redis scan err: %w", err)
	}

	return chatIDs, nil
}
