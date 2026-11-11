package cache

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	gotoCacheKeyPrefix = "short-link:goto:"
	gotoLockKeyPrefix  = "short-link:lock:goto:"
	gotoLockTTL        = 30 * time.Second
	gotoLockRetry      = 50 * time.Millisecond
	gotoNullCacheKeyPrefix = "short-link:is-null:goto:"
	gotoNullCacheValue     = "-"
	gotoNullCacheTTL       = 30 * time.Minute
)

type ShortLinkGotoCache struct {
	client *redis.Client
}

func NewShortLinkGotoCache(client *redis.Client) *ShortLinkGotoCache {
	return &ShortLinkGotoCache{
		client: client,
	}
}

func gotoCacheKey(fullShortURL string) string {
	return gotoCacheKeyPrefix + fullShortURL
}

func gotoLockKey(fullShortURL string) string {
	return gotoLockKeyPrefix + fullShortURL
}

func gotoNullCacheKey(fullShortURL string) string {
	return gotoNullCacheKeyPrefix + fullShortURL
}

func (c *ShortLinkGotoCache) Get(ctx context.Context, fullShortURL string) (string, bool, error) {
	value, err := c.client.Get(ctx, gotoCacheKey(fullShortURL)).Result()
	if err == redis.Nil {
		return "", false, nil
	}
	if err != nil {
		return "", false, fmt.Errorf("读取短链接跳转缓存失败: %w", err)
	}
	return value, true, nil
}

func (c *ShortLinkGotoCache) Set(ctx context.Context, fullShortURL, originURL string) error {
	if err := c.client.Set(ctx, gotoCacheKey(fullShortURL), originURL, 0).Err(); err != nil {
		return fmt.Errorf("写入短链接跳转缓存失败: %w", err)
	}
	return nil
}

func (c *ShortLinkGotoCache) Lock(ctx context.Context, fullShortURL string) (string, error) {
	var value [16]byte
	if _, err := rand.Read(value[:]); err != nil {
		return "", err
	}
	token := hex.EncodeToString(value[:])
	key := gotoLockKey(fullShortURL)

	for {
		acquired, err := c.client.SetNX(ctx, key, token, gotoLockTTL).Result()
		if err != nil {
			return "", fmt.Errorf("获取短链接跳转锁失败: %w", err)
		}
		if acquired {
			return token, nil
		}

		timer := time.NewTimer(gotoLockRetry)
		select {
		case <-ctx.Done():
			timer.Stop()
			return "", ctx.Err()
		case <-timer.C:
		}
	}
}

func (c *ShortLinkGotoCache) Unlock(fullShortURL, token string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	script := `
		if redis.call('GET', KEYS[1]) == ARGV[1] then
			return redis.call('DEL', KEYS[1])
		end
		return 0
	`

	return c.client.Eval(ctx, script, []string{gotoLockKey(fullShortURL)}, token).Err()
}

func (c *ShortLinkGotoCache) GetNull(ctx context.Context, fullShortURL string) (bool, error) {
	value ,err := c.client.Get(ctx, gotoCacheKey(fullShortURL)).Result()
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("读取短链接空缓存失败: %w", err)
	}
	return value != "", nil
}

func (c *ShortLinkGotoCache) SetNull(ctx context.Context, fullShortURL string) error {
	if err := c.client.Set(ctx, gotoNullCacheKey(fullShortURL), gotoNullCacheValue, gotoNullCacheTTL).Err(); err != nil {
		return fmt.Errorf("写入短链接空缓存失败: %w", err)
	}
	return nil
}