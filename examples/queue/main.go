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
	queue := redis.NewRedisQueue(redisClient, "queue-test")

	fmt.Println("starting send message..")
	msgToEnqueue := []string{"aaa", "bbb", "ccc", "ddd"}
	for _, msg := range msgToEnqueue {
		time.Sleep(time.Second)
		fmt.Println("send message - ", msg)
		queue.SendMessage(msg)
	}

	time.Sleep(time.Second * 2)

	fmt.Println("starting receive message..")
	for {
		time.Sleep(time.Second)
		msg, err := queue.ReceiveMessage()
		if err != nil {
			panic(err)
		}
		if msg == "" {
			fmt.Println("receive all messages")
			break
		}
		fmt.Println("receive message - ", msg)
	}

}
