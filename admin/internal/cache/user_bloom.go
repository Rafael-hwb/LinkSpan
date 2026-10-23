package cache

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

const (
	userBloomKey = "linkspan:{user-register}:bloom:v1"
	userReadyKey = "linkspan:{user-register}:ready:v1"
)

type UserBloom struct {
	client *redis.Client
}

func NewUserBloom(client *redis.Client) *UserBloom {
	return &UserBloom{
		client: client,
	}
}

//准备过滤器，标记为尚未就绪（目前为本地单实例启动初始化）
func (b *UserBloom) Prepare(ctx context.Context) error {
	script := `
		redis.call('DEL', KEYS[2])

		if redis.call('EXISTS', KEYS[1]) == 0 then
			redis.call('BF.RESERVE', KEYS[1], ARGV[1], ARGV[2])
		else
			redis.call('BF.INFO', KEYS[1])
		end

		return 1
	`

	err := b.client.Eval(
		ctx,
		script,
		[]string{userBloomKey, userReadyKey},
		0.001,
		100000,
	).Err()
	if err != nil {
		return fmt.Errorf("初始化用户过滤器失败: %w", err)
	}

	return nil
}

//将用户名加入布隆过滤器
func (b *UserBloom) Add(ctx context.Context, username string) error {
	script := `
		if redis.call('EXISTS', KEYS[1]) == 0 then
			return redis.error_reply('user bloom filter missing')
		end

		return redis.call('BF.ADD', KEYS[1], ARGV[1])
	`

	err := b.client.Eval(
		ctx,
		script,
		[]string{userBloomKey},
		username,
	).Err()
	if err != nil {
		return fmt.Errorf("添加用户名到过滤器失败: %w", err)
	}

	return nil
}

//标记已有用户名导入完成
func (b *UserBloom) MarkReady(ctx context.Context) error {
	script := `
		if redis.call('EXISTS', KEYS[1]) == 0 then
			return redis.error_reply('user bloom filter missing')
		end

		redis.call('SET', KEYS[2], '1')
		return 1
	`

	return b.client.Eval(
		ctx,
		script,
		[]string{userBloomKey, userReadyKey},
	).Err()
}

//判断用户名是否可能存在
func (b *UserBloom) Contains(
	ctx context.Context,
	username string,
) (bool, error) {
	script := `
		if redis.call('GET', KEYS[2]) ~= '1' then
			return redis.error_reply('user bloom filter not ready')
		end

		if redis.call('EXISTS', KEYS[1]) == 0 then
			return redis.error_reply('user bloom filter missing')
		end

		return redis.call('BF.EXISTS', KEYS[1], ARGV[1])
	`

	value, err := b.client.Eval(
		ctx,
		script,
		[]string{userBloomKey, userReadyKey},
		username,
	).Int()
	if err != nil {
		return false, fmt.Errorf("检查用户过滤器失败: %w", err)
	}

	return value == 1, nil
}
