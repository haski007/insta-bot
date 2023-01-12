package redis

import (
	"fmt"
)

// keyFromChatID - context is a folder where to store keys to storage diff chats for diff purposes
func (r *redisClient) keyFromChatID(context string, chatID int64) string {
	return fmt.Sprintf("%d/%s", chatID, context)
}

func (r *redisClient) AddChatWithMembers(context string, chatID int64, members []string) error {
	if err := r.conn.LPush(r.keyFromChatID(context, chatID), members).Err(); err != nil {
		return fmt.Errorf("redis LPush err: %w", err)
	}
	return nil
}

func (r *redisClient) ChatExists(context string, chatID int64) (bool, error) {
	exists, err := r.conn.Exists(r.keyFromChatID(context, chatID)).Result()
	if err != nil {
		return false, fmt.Errorf("redis get bool err: %w", err)
	}
	return exists == 1, nil
}

func (r *redisClient) GetChatMembers(context string, chatID int64) (members []string, err error) {

	return r.getStringsArray(r.keyFromChatID(context, chatID))
}

func (r *redisClient) getStringsArray(key string) ([]string, error) {
	arr, err := r.conn.LRange(key, 0, -1).Result()
	if err != nil {
		return nil, fmt.Errorf("redis LRange err: %w", err)
	}

	return arr, nil
}

func (r *redisClient) DeleteChat(context string, chatID int64) error {
	if err := r.conn.Del(r.keyFromChatID(context, chatID)).Err(); err != nil {
		return fmt.Errorf("redis Del err: %w", err)
	}
	return nil
}

// GetAllCTXs - returns map with context as a key and slice of players as a value
func (r *redisClient) GetAllCTXs(chatID int64) (map[string][]string, error) {
	var result = make(map[string][]string)
	iter := r.conn.Scan(0, fmt.Sprintf("%d/*", chatID), 0).Iterator()
	for iter.Next() {
		members, err := r.getStringsArray(iter.Val())
		if err != nil {
			return nil, fmt.Errorf("get strings array err: %w", err)
		}

		if len(members) == 0 {
			continue
		}
		result[iter.Val()] = members
	}

	return result, nil
}
