package main

import (
	"fmt"
	"redis-examples/contract"
	"redis-examples/infra"
	"redis-examples/redis"
	"time"

	"github.com/google/uuid"
)

func main() {
	redisClient, err := infra.ConnectRedis()
	if err != nil {
		panic(err)
	}

	delayedTasks := redis.NewRedisDelayedTask(redisClient)

	items := []contract.Task{
		{UUID: uuid.New().String(), Value: "{\"foo\":\"bar\"}"},
		{UUID: uuid.New().String(), Value: "{\"foo\":\"bar\"}"},
	}

	for _, item := range items {
		delayedTasks.AddDelayed(item, time.Second*5)
	}

	// should print 0 since none is ready
	tasks := delayedTasks.GetReadyTasks()
	fmt.Println("Received tasks:", tasks)

	time.Sleep(time.Second * 6)

	// should print 2 tasks
	tasks = delayedTasks.GetReadyTasks()
	fmt.Println("Received tasks:", tasks)
}
