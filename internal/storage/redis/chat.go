package redis

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/haski007/insta-bot/internal/storage"
)

// keyFromChatID - context is a folder where to store keys to storage diff chats for diff purposes
func (r *redisClient) keyFromChatID(context string, chatID int64) string {
	return fmt.Sprintf("%d/%s", chatID, context)
}

// keyFromChatIDPollID
func (r *redisClient) keyFromChatIDPollID(context string, chatID int64, pollID string) string {
	return fmt.Sprintf("%s/%s", r.keyFromChatID(context, chatID), pollID)
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

func getStartupSubKey(chatID int64) string {
	return fmt.Sprintf("sub/startup/%d", chatID)
}

func (r *redisClient) SubscribeChatToStartup(chatID int64) error {
	key := getStartupSubKey(chatID)

	if err := r.conn.LPush(key, storage.Replica{
		Role: "user",
		Content: `Your task in this chat will be to generate 3 new startup ideas everytime I ask you to do it with a phrase "Generate new startup ideas",
startup ideas should be not to complex it should be possible to implement them in 1-2 months for 1-2 developers on golang`,
	}).Err(); err != nil {
		return fmt.Errorf("redis lpush err: %w", err)
	}
	return nil
}

func (r *redisClient) UnsubscribeChatToStartup(chatID int64) error {
	key := getStartupSubKey(chatID)

	if err := r.conn.Del(key).Err(); err != nil {
		return fmt.Errorf("redis del err: %w", err)
	}
	return nil
}

func (r *redisClient) PushStartupNewsletter(chatID int64, replicas []storage.Replica) error {
	key := getStartupSubKey(chatID)
	for _, replica := range replicas {
		if err := r.conn.LPush(key, replica).Err(); err != nil {
			return fmt.Errorf("redis lpush err: %w", err)
		}
	}

	return nil
}

func (r *redisClient) GetStartupNewsletter(chatID int64) (conversation []storage.Replica, err error) {
	key := getStartupSubKey(chatID)
	return r.getReplicasArray(key)
}

func (r *redisClient) GetAllStartupNewsletters() (newsletterConv map[int64][]storage.Replica, err error) {
	newsletterConv = make(map[int64][]storage.Replica)
	iter := r.conn.Scan(0, "sub/startup/*", 0).Iterator()
	for iter.Next() {
		chatID, err := strconv.ParseInt(strings.Split(iter.Val(), "/")[2], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("parse int err: %w", err)
		}

		newsletter, err := r.GetStartupNewsletter(chatID)
		if err != nil {
			return nil, fmt.Errorf("get pull err: %w", err)
		}

		newsletterConv[chatID] = newsletter
	}

	return newsletterConv, nil
}

func (r *redisClient) EnableLoaderForChat(chatID int64) error {
	return r.conn.SRem(r.keyFromChatID("loader_disabled", chatID), chatID).Err()
}

func (r *redisClient) DisableLoaderForChat(chatID int64) error {
	return r.conn.SAdd(r.keyFromChatID("loader_disabled", chatID), chatID).Err()
}

func (r *redisClient) IsChatLoaderEnabled(chatID int64) (bool, error) {
	exists, err := r.conn.SIsMember(r.keyFromChatID("loader_disabled", chatID), chatID).Result()
	if err != nil {
		return false, fmt.Errorf("redis sismember err: %w", err)
	}
	return !exists, nil
}
