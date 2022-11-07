package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type RedisPubSub struct {
	channel string
	redis   *redis.Client
}

func NewRedisPubSub(redis *redis.Client, channel string) *RedisPubSub {
	return &RedisPubSub{redis: redis, channel: channel}
}

func (r *RedisPubSub) SendMessage(msg string) (err error) {
	return r.redis.Publish(context.Background(), r.channel, msg).Err()
}

func (r RedisPubSub) ReceiveMessage(ctx context.Context, chanMessage chan string) (err error) {
	subscriber := r.redis.Subscribe(context.Background(), r.channel)
loop:
	for {
		select {
		case <-ctx.Done():
			break loop

		default:
			msg, err := subscriber.ReceiveMessage(context.Background())
			if err != nil {
				panic(err)
			}
			if msg.String() != "" {
				chanMessage <- msg.Payload
			}
		}

	}
	return nil
}
