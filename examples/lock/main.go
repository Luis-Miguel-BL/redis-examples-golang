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
	lock_prefix := "test-lock"
	lock_key_test := "lock1"

	lock := redis.NewRedisLock(lock_prefix, redisClient)
	success, _ := lock.Set(lock_key_test, time.Second)
	fmt.Printf("\n First Lock: %t \n", success)

	locked := lock.Get(lock_key_test)
	fmt.Println("locked: ", locked)

	success, _ = lock.Set(lock_key_test, time.Second)
	fmt.Printf("\n Second Lock: %t \n", success)

	time.Sleep(time.Second)

	lockedAfterSleep := lock.Get(lock_key_test)
	fmt.Println("lockedAfterSleep: ", lockedAfterSleep)

	success, _ = lock.Set(lock_key_test, time.Second)
	fmt.Printf("\n Third Lock: %t \n", success)

	lockedAfterThirdLock := lock.Get(lock_key_test)
	fmt.Println("lockedAfterThirdLock: ", lockedAfterThirdLock)
}
