package redis

import (
	"fmt"
)

const ukraineForUkrainiansSubKey = "ukraine_for_ukrainians"

func (r *redisClient) SubscribeChatToUkraineForUkrainians(chatID int64) error {
	return r.conn.SAdd(r.keyFromChatID(ukraineForUkrainiansSubKey, chatID), chatID).Err()
}

func (r *redisClient) UnsubscribeChatFromUkraineForUkrainians(chatID int64) error {
	if err := r.conn.SRem(r.keyFromChatID(ukraineForUkrainiansSubKey, chatID), chatID).Err(); err != nil {
		return fmt.Errorf("redis SRem ukr sub: %w", err)
	}
	_ = r.UkraineAnglicismIgnorePurgeChat(chatID)
	return nil
}

func (r *redisClient) IsChatSubscribedToUkraineForUkrainians(chatID int64) (bool, error) {
	exists, err := r.conn.SIsMember(r.keyFromChatID(ukraineForUkrainiansSubKey, chatID), chatID).Result()
	if err != nil {
		return false, fmt.Errorf("redis sismember err: %w", err)
	}
	return exists, nil
}
