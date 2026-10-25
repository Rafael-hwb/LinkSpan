package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

const UserLoginKeyPredix = "short-link:login:"

type UserSession struct {
	client *redis.Client
}

func NewUserSession(client *redis.Client) *UserSession {
	return &UserSession{
		client:  client,
	}
}

func (s *UserSession) Exists(ctx context.Context, username string, 
	token string) (bool, error) {
		return s.client.HExists(ctx, UserLoginKeyPredix+username, 
			token).Result()
}
