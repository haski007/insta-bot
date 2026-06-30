package redis

import "fmt"

const fuckListCtx = "fuck_list"

func (r *redisClient) FuckListAdd(username string) error {
	if err := r.conn.SAdd(r.keyFromChatID(fuckListCtx, 0), username).Err(); err != nil {
		return fmt.Errorf("redis SAdd fuck list: %w", err)
	}
	return nil
}

func (r *redisClient) FuckListRemove(username string) error {
	if err := r.conn.SRem(r.keyFromChatID(fuckListCtx, 0), username).Err(); err != nil {
		return fmt.Errorf("redis SRem fuck list: %w", err)
	}
	return nil
}

func (r *redisClient) FuckListContains(username string) (bool, error) {
	ok, err := r.conn.SIsMember(r.keyFromChatID(fuckListCtx, 0), username).Result()
	if err != nil {
		return false, fmt.Errorf("redis SIsMember fuck list: %w", err)
	}
	return ok, nil
}
