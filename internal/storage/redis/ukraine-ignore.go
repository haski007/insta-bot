package redis

import (
	"fmt"
)

const ukraineAnglicismIgnoreCtx = "ukr_ign"

func (r *redisClient) UkraineAnglicismIgnoreAdd(chatID int64, username string) error {
	if err := r.conn.SAdd(r.keyFromChatID(ukraineAnglicismIgnoreCtx, chatID), username).Err(); err != nil {
		return fmt.Errorf("redis SAdd ukr ignore: %w", err)
	}
	return nil
}

func (r *redisClient) UkraineAnglicismIgnoreRemove(chatID int64, username string) error {
	if err := r.conn.SRem(r.keyFromChatID(ukraineAnglicismIgnoreCtx, chatID), username).Err(); err != nil {
		return fmt.Errorf("redis SRem ukr ignore: %w", err)
	}
	return nil
}

func (r *redisClient) UkraineAnglicismIgnoreContains(chatID int64, username string) (bool, error) {
	ok, err := r.conn.SIsMember(r.keyFromChatID(ukraineAnglicismIgnoreCtx, chatID), username).Result()
	if err != nil {
		return false, fmt.Errorf("redis SIsMember ukr ignore: %w", err)
	}
	return ok, nil
}

// UkraineAnglicismIgnorePurgeChat removes all ignored nicknames for a chat (e.g. on unsub).
func (r *redisClient) UkraineAnglicismIgnorePurgeChat(chatID int64) error {
	if err := r.conn.Del(r.keyFromChatID(ukraineAnglicismIgnoreCtx, chatID)).Err(); err != nil {
		return fmt.Errorf("redis Del ukr ignore: %w", err)
	}
	return nil
}
