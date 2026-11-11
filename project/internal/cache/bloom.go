package cache

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

const (
	shortLinkBloomKey = "short-link:bloom:full-short-url"
	shortLinkReadyKey = "short-link:bloom:full-short-url:ready"
)

type ShortLinkBloom struct {
	client *redis.Client
}

func NewShortLinkBloom(client *redis.Client) *ShortLinkBloom {
	return &ShortLinkBloom{client: client}
}

func (b *ShortLinkBloom) Prepare(ctx context.Context) error {
	script := `
		redis.call('DEL', KEYS[2])
		if redis.call('EXISTS', KEYS[1]) == 0 then
			redis.call('BF.RESERVE', KEYS[1], ARGV[1], ARGV[2])
		else
			redis.call('BF.INFO', KEYS[1])
		end
		return 1
	`
	if err := b.client.Eval(ctx, script, []string{shortLinkBloomKey, shortLinkReadyKey}, 0.001, 100000).Err(); err != nil {
		return fmt.Errorf("初始化短链接过滤器失败: %w", err)
	}
	return nil
}

func (b *ShortLinkBloom) Add(ctx context.Context, fullShortURL string) error {
	script := `
		if redis.call('EXISTS', KEYS[1]) == 0 then
			return redis.error_reply('short link bloom filter missing')
		end
		return redis.call('BF.ADD', KEYS[1], ARGV[1])
	`
	if err := b.client.Eval(ctx, script, []string{shortLinkBloomKey}, fullShortURL).Err(); err != nil {
		return fmt.Errorf("添加短链接到过滤器失败: %w", err)
	}
	return nil
}

func (b *ShortLinkBloom) MarkReady(ctx context.Context) error {
	script := `
		if redis.call('EXISTS', KEYS[1]) == 0 then
			return redis.error_reply('short link bloom filter missing')
		end
		redis.call('SET', KEYS[2], '1')
		return 1
	`
	if err := b.client.Eval(ctx, script, []string{shortLinkBloomKey, shortLinkReadyKey}).Err(); err != nil {
		return fmt.Errorf("标记短链接过滤器就绪失败: %w", err)
	}
	return nil
}

func (b *ShortLinkBloom) Contains(ctx context.Context, fullShortURL string) (bool, error) {
	script := `
		if redis.call('GET', KEYS[2]) ~= '1' then
			return redis.error_reply('short link bloom filter not ready')
		end
		if redis.call('EXISTS', KEYS[1]) == 0 then
			return redis.error_reply('short link bloom filter missing')
		end
		return redis.call('BF.EXISTS', KEYS[1], ARGV[1])
	`
	value, err := b.client.Eval(ctx, script, []string{shortLinkBloomKey, shortLinkReadyKey}, fullShortURL).Int()
	if err != nil {
		return false, fmt.Errorf("检查短链接过滤器失败: %w", err)
	}
	return value == 1, nil
}
