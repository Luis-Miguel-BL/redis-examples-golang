package main

import (
	"fmt"
	"redis-examples/infra"
	"redis-examples/redis"
	"time"
)

func main() {
	redisClient, err := infra.ConnectRedis()
	if err != nil {
		panic(err)
	}
	semaphoreOptions := redis.RedisSemaphoreOptions{
		Client:               redisClient,
		Key:                  "semaphore-test",
		MaxParallelResources: 5,
		LockExpiration:       time.Second,
	}
	semaphore := redis.NewRedisSemaphore(&semaphoreOptions)

	for true {
		err := semaphore.Lock()
		if err != nil {
			// limit reached
			continue
		}
		processSomething()
	}
}

func processSomething() {
	fmt.Print(".")
	time.Sleep(time.Millisecond * 100)
}
