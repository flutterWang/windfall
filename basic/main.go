package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456",
		DB:       0,
	})

	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	err = client.Set("key", "newValue", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err = client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}
