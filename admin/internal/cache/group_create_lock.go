package cache

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/redis/go-redis/v9"
)

type GroupCreateLock struct {
	client *redis.Client
}

func NewGroupCreateLock(client *redis.Client) *GroupCreateLock {
	return &GroupCreateLock{client: client}
}

func groupCreateLockKey(username string) string {
	return "linkspan:group-create:lock:" + username
}

// TryLock 使用固定 30 秒租期；不提供 Redisson 的自动续期能力。
func (l *GroupCreateLock) TryLock(ctx context.Context, username string) (string, bool, error) {
	var value [16]byte
	if _, err := rand.Read(value[:]); err != nil {
		return "", false, err
	}
	token := hex.EncodeToString(value[:])
	acquired, err := l.client.SetNX(ctx, groupCreateLockKey(username), token, 30*time.Second).Result()
	if err != nil {
		return "", false, err
	}
	return token, acquired, nil
}

func (l *GroupCreateLock) Unlock(ctx context.Context, username, token string) error {
	script := `
		if redis.call('GET', KEYS[1]) == ARGV[1] then
			return redis.call('DEL', KEYS[1])
		end
		return 0
	`
	return l.client.Eval(ctx, script, []string{groupCreateLockKey(username)}, token).Err()
}
