package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func getHistoryMessageKey(chatID int64, messageID int) string {
	return fmt.Sprintf("message/%d/%d", chatID, messageID)
}

func getChatHistoryKey(chatID int64) string {
	return fmt.Sprintf("message/%d/*", chatID)
}

func getHistoryIndexKey(chatID int64) string {
	return fmt.Sprintf("history_index/%d", chatID)
}

func (r *redisClient) SaveMessage(chatID int64, messageID int, message string) error {
	key := getHistoryMessageKey(chatID, messageID)

	if err := r.conn.Set(key, message, r.historyTTL).Err(); err != nil {
		return fmt.Errorf("redis set err: %w", err)
	}
	if err := r.conn.ZAdd(getHistoryIndexKey(chatID), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: key,
	}).Err(); err != nil {
		return fmt.Errorf("redis set err: %w", err)
	}

	return nil
}

func (r *redisClient) GetMessages(chatID int64, count int) (messages []string, err error) {
	messagesKeys, err := r.conn.ZRevRange(getHistoryIndexKey(chatID), 0, int64(count-1)).Result()
	if err != nil {
		return nil, fmt.Errorf("redis zrevrange err: %w", err)
	}
	// reverse messagesKeys
	for i, j := 0, len(messagesKeys)-1; i < j; i, j = i+1, j-1 {
		messagesKeys[i], messagesKeys[j] = messagesKeys[j], messagesKeys[i]
	}

	res, err := r.conn.MGet(messagesKeys...).Result()
	if err != nil {
		return nil, fmt.Errorf("redis mget err: %w", err)
	}
	messages = make([]string, 0, len(res))
	for _, v := range res {
		messages = append(messages, v.(string))
	}

	return messages, nil
}

func (r *redisClient) PurgeHistory(chatID int64) error {
	if err := r.conn.Del(getChatHistoryKey(chatID)).Err(); err != nil {
		return fmt.Errorf("redis del chat history err: %w", err)
	}

	if err := r.conn.Del(getHistoryIndexKey(chatID)).Err(); err != nil {
		return fmt.Errorf("redis del history index for chat err: %w", err)
	}

	// Find keys matching the pattern
	var cursor uint64
	var err error
	var keys []string

	for {
		// Scan with cursor
		var scanKeys []string
		scanKeys, cursor, err = r.conn.Scan(cursor, getChatHistoryKey(chatID), 0).Result()
		if err != nil {
			return fmt.Errorf("redis scan err: %w", err)
		}

		keys = append(keys, scanKeys...)

		if cursor == 0 {
			break
		}
	}

	if len(keys) > 0 {
		err = r.conn.Del(keys...).Err()
		if err != nil {
			return fmt.Errorf("redis del err: %w", err)
		}
	}

	return nil
}
