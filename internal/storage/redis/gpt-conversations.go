package redis

import (
	"errors"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/haski007/insta-bot/internal/storage"
)

func getConversationKey(username string, chatID, userID int64) string {
	return fmt.Sprintf("gpt/conversations/%s:%d:%d", username, chatID, userID)
}

func getSystemRoleKey(chatID int64) string {
	return fmt.Sprintf("gpt/system-role/%d", chatID)
}

func (r *redisClient) SetSystemRoleForChat(chatID int64, role string) error {
	key := getSystemRoleKey(chatID)
	if err := r.conn.Set(key, role, 0).Err(); err != nil {
		return fmt.Errorf("redis set err: %w", err)
	}
	return nil
}

func (r *redisClient) GetSystemRole(chatID int64) (role string, err error) {
	key := getSystemRoleKey(chatID)
	if err := r.conn.Get(key).Scan(&role); err != nil {
		if errors.Is(err, redis.Nil) {
			return "", storage.ErrNotFound
		}
		return "", fmt.Errorf("redis get err: %w", err)
	}
	return role, nil
}

func (r *redisClient) DropConversation(req *storage.DropConversationReq) (err error) {
	key := getConversationKey(req.Username, req.ChatID, req.UserID)
	if err := r.conn.Del(key).Err(); err != nil {
		return fmt.Errorf("redis del err: %w", err)
	}
	return nil
}

func (r *redisClient) PushConversation(req *storage.PushConversationReq) (err error) {
	key := getConversationKey(req.Username, req.ChatID, req.UserID)
	for _, replica := range req.Replicas {
		if err := r.conn.LPush(key, replica).Err(); err != nil {
			return fmt.Errorf("redis lpush err: %w", err)
		}
	}

	if err := r.conn.Expire(key, r.convTTL).Err(); err != nil {
		return fmt.Errorf("redis expire err: %w", err)
	}

	return nil
}

func (r *redisClient) GetConversation(req *storage.GetConversationReq) (conversation []storage.Replica, err error) {
	key := getConversationKey(req.Username, req.ChatID, req.UserID)

	return r.getReplicasArray(key)
}

func (r *redisClient) getReplicasArray(key string) (conversation []storage.Replica, err error) {
	err = r.conn.LRange(key, 0, -1).ScanSlice(&conversation)
	if err != nil {
		return nil, fmt.Errorf("redis LRange err: %w", err)
	}

	return conversation, nil
}
