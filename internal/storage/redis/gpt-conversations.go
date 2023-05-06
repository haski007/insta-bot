package redis

import (
	"fmt"

	"github.com/haski007/insta-bot/internal/storage"
)

func getConversationKey(username string, chatID, userID int64) string {
	return fmt.Sprintf("conversations/%s:%d:%d", username, chatID, userID)
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
