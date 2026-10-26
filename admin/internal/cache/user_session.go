package cache

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

const UserLoginKeyPrefix = "short-link:login:"
const userSessionTTL = 30 * time.Minute

type SessionUser struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	RealName string `json:"realName"`
}

type UserSession struct {
	client *redis.Client
}

func NewUserSession(client *redis.Client) *UserSession {
	return &UserSession{
		client: client,
	}
}

func (s *UserSession) Exists(ctx context.Context, username string,
	token string) (bool, error) {
	return s.client.HExists(ctx, UserLoginKeyPrefix+username,
		token).Result()
}

func (s *UserSession) CreateOrRefresh(ctx context.Context, user SessionUser) (string, error) {
	payload, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	var random [32]byte
	if _, err := rand.Read(random[:]); err != nil {
		return "", err
	}

	candidateToken := hex.EncodeToString(random[:])

	script := `
		local tokens = redis.call('HKEYS', KEYS[1])
		local token = ARGV[1]

		if #tokens > 0 then
			token = tokens[1]
		end

		redis.call('HSET', KEYS[1], token, ARGV[2])
		redis.call('EXPIRE', KEYS[1], ARGV[3])

		return token
	`

	return s.client.Eval(
		ctx,
		script,
		[]string{UserLoginKeyPrefix + user.Username},
		candidateToken,
		string(payload),
		int64(userSessionTTL/time.Second),
	).Text()
}

func (s *UserSession) DeleteIfValid(ctx context.Context, username string, token string) (bool, error) {
	script := `
		if redis.call('HEXISTS', KEYS[1], ARGV[1]) == 1 then
			return redis.call('DEL', KEYS[1])
		end

		return 0
	`

	deleted, err := s.client.Eval(ctx, script, []string{UserLoginKeyPrefix + username}, token).Int()

	if err != nil {
		return false, err
	}

	return deleted == 1, nil
}
