package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"redis-examples/contract"
	"redis-examples/infra"
	"redis-examples/redis"
	"time"
)

var timeLastRequest time.Time
var redisCache contract.Cache

func hello(w http.ResponseWriter, req *http.Request) {
	timeLastRequest := time.Now()
	responseData := map[string]string{
		"time_last_request": timeLastRequest.String(),
	}

	time.Sleep(time.Millisecond * 200)

	err := redisCache.Set(req.URL.Path, responseData, time.Second*10)
	if err != nil {
		panic(err)
	}

	responseBytes, err := json.Marshal(responseData)
	if err != nil {
		panic(err)
	}
	_, err = w.Write(responseBytes)
	if err != nil {
		panic(err)
	}
}

func applicationJSON(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	}
}

func cacheRoute(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, _ := redisCache.Get(r.URL.Path)
		if data == nil {
			h.ServeHTTP(w, r)
			return
		}

		dataBytes, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		_, err = w.Write(dataBytes)
		if err != nil {
			panic(err)
		}

	}
}

func main() {
	redisClient, err := infra.ConnectRedis()
	if err != nil {
		panic(err)
	}
	redisCache = redis.NewRedisCache(redisClient)

	http.HandleFunc("/hello", applicationJSON(cacheRoute(hello)))

	fmt.Println("Runninng server at localhost:", "8090")
	http.ListenAndServe(":8090", nil)

}
