package infra

import (
	"context"
	"errors"
	"log"

	"github.com/go-redis/redis/v8"
)

const (
	successPing = "PONG"
	errConnect  = "cannot be connect to redis"
)

func ConnectRedis() (redisClient *redis.Client, err error) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       2,
		Password: "",
	})

	log.Println("connecting to redis...")

	if successPing != redisClient.Ping(context.Background()).Val() {
		return redisClient, errors.New(errConnect)
	}

	log.Println("sucessfuly connected to redis!")

	return redisClient, nil
}
