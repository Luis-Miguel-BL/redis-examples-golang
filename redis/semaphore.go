package redis

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisSemaphoreOptions struct {
	Client               *redis.Client
	LockExpiration       time.Duration
	Key                  string
	MaxParallelResources int
}

type RedisSemaphore struct {
	client               *redis.Client
	Key                  string
	LockExpiration       time.Duration
	MaxParallelResources int
	LockScript           string
}

const (
	lockCommand = `
		local quantityLocks = redis.call('INCR', KEYS[1]);
		if (quantityLocks == 1)  then
			redis.call('EXPIRE', KEYS[1], ARGV[2])
		elseif quantityLocks > tonumber(ARGV[1]) then
			redis.call('DECR', KEYS[1])
			return 0
		end
		return 1`

	lockSuccess int64 = 1
)

func NewRedisSemaphore(options *RedisSemaphoreOptions) *RedisSemaphore {
	lockScriptLoaded := options.Client.ScriptLoad(context.Background(), lockCommand)

	return &RedisSemaphore{
		client:               options.Client,
		Key:                  options.Key,
		LockExpiration:       options.LockExpiration,
		MaxParallelResources: options.MaxParallelResources,
		LockScript:           lockScriptLoaded.Val(),
	}
}

func (r *RedisSemaphore) Lock() (err error) {
	scriptKeys := []string{r.Key}
	scriptArgs := []string{
		fmt.Sprint(r.MaxParallelResources),
		fmt.Sprint(r.LockExpiration.Seconds()),
	}
	resp, err := r.client.EvalSha(context.Background(), r.LockScript, scriptKeys, scriptArgs).Result()
	if err != nil {
		return err
	}
	if resp != lockSuccess {
		return errors.New("lock limit reached")
	}

	return nil
}
