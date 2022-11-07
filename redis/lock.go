package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisLock struct {
	KeyPrefix string
	redis     *redis.Client
}

func NewRedisLock(keyPrefix string, redis *redis.Client) *RedisLock {
	return &RedisLock{KeyPrefix: keyPrefix, redis: redis}
}

func (l *RedisLock) Get(key string) (hasLock bool) {
	keyWithPrefix := fmt.Sprintf("%s-%s", l.KeyPrefix, key)
	alreadySent := l.redis.Exists(context.Background(), keyWithPrefix).Val() > 0
	return alreadySent
}
func (l *RedisLock) Set(key string, lockExpiration time.Duration) (success bool, err error) {
	keyWithPrefix := fmt.Sprintf("%s-%s", l.KeyPrefix, key)
	success, err = l.redis.SetNX(context.Background(), keyWithPrefix, time.Now().Unix(), lockExpiration).Result()
	return success, err
}
