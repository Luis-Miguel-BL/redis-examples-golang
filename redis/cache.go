package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	redis *redis.Client
}

func NewRedisCache(redis *redis.Client) *RedisCache {
	return &RedisCache{redis: redis}
}

func (r *RedisCache) Get(route string) (data interface{}, err error) {
	redisData, err := r.redis.Get(context.Background(), route).Result()
	if err != nil {
		return data, err
	}

	err = json.Unmarshal([]byte(redisData), &data)
	return data, err
}

func (r *RedisCache) Set(route string, data interface{}, expiration time.Duration) (err error) {
	dataString, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return r.redis.Set(context.Background(), route, string(dataString), expiration).Err()
}
