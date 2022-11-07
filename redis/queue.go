package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type RedisQueue struct {
	channel string
	redis   *redis.Client
}

func NewRedisQueue(redis *redis.Client, channel string) *RedisQueue {
	return &RedisQueue{redis: redis, channel: channel}
}

func (r *RedisQueue) SendMessage(msg string) (err error) {
	return r.redis.RPush(context.Background(), r.channel, msg).Err()
}

func (r RedisQueue) ReceiveMessage() (msg string, err error) {
	msg, err = r.redis.LPop(context.Background(), r.channel).Result()

	if err != nil && err != redis.Nil {
		return "", err
	}

	if len(msg) == 0 {
		return "", nil
	}

	return msg, nil
}
