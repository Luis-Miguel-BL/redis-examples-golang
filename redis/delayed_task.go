package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"redis-examples/contract"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisDelayedTask struct {
	redis *redis.Client
}

const delayedList = "delayed"

func NewRedisDelayedTask(redis *redis.Client) *RedisDelayedTask {
	return &RedisDelayedTask{redis: redis}
}

func (d *RedisDelayedTask) AddDelayed(value contract.Task, delay time.Duration) {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	taskReadyInSeconds := time.Now().Add(delay).Unix()
	member := redis.Z{
		Score:  float64(taskReadyInSeconds),
		Member: jsonValue,
	}
	_, err = d.redis.ZAdd(context.Background(), delayedList, &member).Result()
	if err != nil {
		fmt.Println(err)
	}
}

func (d *RedisDelayedTask) GetReadyTasks() []contract.Task {
	maxTime := time.Now().Unix()
	opt := redis.ZRangeBy{
		Min: fmt.Sprintf("%d", 0),
		Max: fmt.Sprintf("%d", maxTime),
	}
	cmd := d.redis.ZRevRangeByScore(context.Background(), delayedList, &opt)
	resultSet, err := cmd.Result()
	if err != nil {
		fmt.Println(err)
		panic("redis_error")
	}

	tasks := make([]contract.Task, len(resultSet))
	for i, t := range resultSet {
		err := json.Unmarshal([]byte(t), &tasks[i])
		if err != nil {
			fmt.Println(err)
			panic("JSON!!!")
		}
	}

	d.RemoveTasks(tasks)
	return tasks
}

func (d *RedisDelayedTask) RemoveTasks(tasks []contract.Task) {
	fmt.Println("ZRem called with: ", len(tasks), "tasks")
	if len(tasks) == 0 {
		return
	}

	jsonTasks := make([]string, len(tasks))
	for i, t := range tasks {
		jsonValue, err := json.Marshal(t)
		if err != nil {
			panic(err)
		}
		jsonTasks[i] = string(jsonValue)
	}

	_, err := d.redis.ZRem(context.Background(), delayedList, jsonTasks).Result()
	if err != nil {
		fmt.Println(err)
		panic("redis_error")
	}
}
