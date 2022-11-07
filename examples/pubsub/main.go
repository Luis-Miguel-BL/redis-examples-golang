package main

import (
	"context"
	"fmt"
	"redis-examples/infra"
	"redis-examples/redis"
	"sync"
	"time"
)

func main() {
	ctx, stop := context.WithCancel(context.Background())
	var wg = sync.WaitGroup{}
	redisClient, err := infra.ConnectRedis()
	if err != nil {
		panic(err)
	}
	pubsub := redis.NewRedisPubSub(redisClient, "pubsub-test")

	err = pubsub.SendMessage("message-lost")
	if err != nil {
		panic(err)
	}

	chanMessage := make(chan string)
	wg.Add(1)
	go func() {
		defer wg.Done()
		pubsub.ReceiveMessage(ctx, chanMessage)
	}()

	wg.Add(1)
	go func(ctx context.Context) {
	loop:
		for {
			select {
			case <-ctx.Done():
				wg.Done()
				break loop
			default:
				fmt.Println("message received: ", <-chanMessage)
			}
		}
	}(ctx)

	time.Sleep(time.Second)

	for i := 0; i < 10; i++ {
		pubsub.SendMessage(fmt.Sprintf("message-%d", i))
	}
	stop()
	wg.Wait()
	fmt.Println("fim")
}
