package util

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisLock struct {
	client   *redis.ClusterClient
	lockKey  string
	timeout  time.Duration
	needWait bool
}

func NewRedisLock(client *redis.ClusterClient, lockKey string, timeout time.Duration, needWait bool) *RedisLock {
	return &RedisLock{
		client:   client,
		lockKey:  lockKey,
		timeout:  timeout,
		needWait: needWait,
	}
}

func (r *RedisLock) Lock(ctx context.Context) bool {
	isLocked := r.client.SetNX(ctx, r.lockKey, 1, r.timeout).Val()
	count := 0
	for !isLocked && r.needWait {
		if count > 100 {
			break
		}
		isLocked = r.client.SetNX(ctx, r.lockKey, 1, r.timeout).Val()
		if !isLocked {
			count++
			time.Sleep(100 * time.Millisecond)
		}
	}
	return isLocked
}

func (r *RedisLock) Unlock(ctx context.Context) error {
	_, err := r.client.Del(ctx, r.lockKey).Result()
	return err
}
