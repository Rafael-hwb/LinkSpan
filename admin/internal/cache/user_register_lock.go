package cache

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/redis/go-redis/v9"
)

type UserRegisterLock struct {
	client *redis.Client
}

func NewUserRegisterLock(client *redis.Client) *UserRegisterLock {
	return &UserRegisterLock{
		client: client,
	}
}

func registerLockKey(username string) string {
	return "linkspan:user-register:lock:" + username
}

//尝试取得锁，不等待其他持有者释放
func (l *UserRegisterLock) TryLock(ctx context.Context, 
	username string) (string, bool, error) {
		var random [16]byte

		if _, err := rand.Read(random[:]); err != nil {
			return "", false, err
		}

		token := hex.EncodeToString(random[:])

		acquired, err := l.client.SetNX(
			ctx,
			registerLockKey(username),
			token,
			30 * time.Second,
		).Result()

		if err != nil {
			return "", false, err
		}

		return token, acquired, nil
	}


//删除token仍属于当前请求的锁
func (l *UserRegisterLock) Unlock(
	ctx context.Context,
	username string,
	token string,
) error {
	script := `
		if redis.call('GET', KEYS[1]) == ARGV[1] then
			return redis.call('DEL', KEYS[1])
		end

		return 0
	`

	return l.client.Eval(
		ctx,
		script,
		[]string{registerLockKey(username)},
		token,
	).Err()
}