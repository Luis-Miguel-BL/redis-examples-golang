package main

import (
	"context"
	"fmt"
	"os"
	"redis-examples/infra"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	redisClient, err := infra.ConnectRedis()
	if err != nil {
		panic(err)
	}

	// Comando para o redis começar a publicar os eventos
	_, err = redisClient.Do(context.Background(), "CONFIG", "SET", "notify-keyspace-events", "KEA").Result()
	if err != nil {
		fmt.Printf("unable to set keyspace events %v", err.Error())
		os.Exit(1)
	}

	// Inscrevendo no canal para receber os eventos de 'expired' publicados pelo redis no db '2'
	pubsub := redisClient.PSubscribe(context.Background(), "__keyevent@2__:expired")
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func(redis.PubSub) {
		exitLoopCounter := 0
		for {
			message, err := pubsub.ReceiveMessage(context.Background())
			exitLoopCounter++
			if err != nil {
				fmt.Printf("error message - %v", err.Error())
				break
			}

			fmt.Printf("Keyspace event recieved %v \n", message.String())
			if exitLoopCounter == 10 {
				wg.Done()
			}
		}
	}(*pubsub)

	wg.Add(1)
	go func(redis.Client, *sync.WaitGroup) {
		for i := 0; i <= 10; i++ {
			dynamicKey := fmt.Sprintf("event_%v", i)
			redisClient.Set(context.Background(), dynamicKey, "someval", time.Second).Result()
		}
		wg.Done()
	}(*redisClient, wg)

	wg.Wait()
}
