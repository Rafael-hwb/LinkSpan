package cache

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

const (
	groupBloomKey = "linkspan:{group-register}:bloom:v1"
	groupReadyKey = "linkspan:{group-register}:ready:v1"
)

type GroupBloom struct {
	client *redis.Client
}

func NewGroupBloom(client *redis.Client) *GroupBloom {
	return &GroupBloom{client: client}
}

// Prepare 仅用于本地单实例启动：保留已有过滤器，导入完成前拒绝业务检查。
func (b *GroupBloom) Prepare(ctx context.Context) error {
	script := `
		redis.call('DEL', KEYS[2])
		if redis.call('EXISTS', KEYS[1]) == 0 then
			redis.call('BF.RESERVE', KEYS[1], ARGV[1], ARGV[2])
		else
			redis.call('BF.INFO', KEYS[1])
		end
		return 1
	`
	if err := b.client.Eval(ctx, script, []string{groupBloomKey, groupReadyKey}, 0.001, 100000).Err(); err != nil {
		return fmt.Errorf("初始化分组过滤器失败: %w", err)
	}
	return nil
}

func (b *GroupBloom) Add(ctx context.Context, gid string) error {
	script := `
		if redis.call('EXISTS', KEYS[1]) == 0 then
			return redis.error_reply('group bloom filter missing')
		end
		return redis.call('BF.ADD', KEYS[1], ARGV[1])
	`
	if err := b.client.Eval(ctx, script, []string{groupBloomKey}, gid).Err(); err != nil {
		return fmt.Errorf("添加分组标识到过滤器失败: %w", err)
	}
	return nil
}

func (b *GroupBloom) MarkReady(ctx context.Context) error {
	script := `
		if redis.call('EXISTS', KEYS[1]) == 0 then
			return redis.error_reply('group bloom filter missing')
		end
		redis.call('SET', KEYS[2], '1')
		return 1
	`
	if err := b.client.Eval(ctx, script, []string{groupBloomKey, groupReadyKey}).Err(); err != nil {
		return fmt.Errorf("标记分组过滤器就绪失败: %w", err)
	}
	return nil
}

func (b *GroupBloom) Contains(ctx context.Context, gid string) (bool, error) {
	script := `
		if redis.call('GET', KEYS[2]) ~= '1' then
			return redis.error_reply('group bloom filter not ready')
		end
		if redis.call('EXISTS', KEYS[1]) == 0 then
			return redis.error_reply('group bloom filter missing')
		end
		return redis.call('BF.EXISTS', KEYS[1], ARGV[1])
	`
	value, err := b.client.Eval(ctx, script, []string{groupBloomKey, groupReadyKey}, gid).Int()
	if err != nil {
		return false, fmt.Errorf("检查分组过滤器失败: %w", err)
	}
	return value == 1, nil
}
