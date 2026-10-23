package bootstrap

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

func OpenRedis() (*redis.Client, error) {
	addr := os.Getenv("LINKSPAN_REDIS_ADDR")

	if addr == "" {
		addr = "127.0.0.1:6380"
	}

	client := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     os.Getenv("LINKSPAN_REDIS_PASSWORD"),
		DB:           0,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	})

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		client.Close()
		return nil, fmt.Errorf("连接 Redis 失败: %w", err)
	}

	return client, nil
}
