package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456",
		DB:       0,
	})

	pipe := client.Pipeline()

	incr := pipe.Incr("bofeng")
	pipe.Expire("bofeng", time.Hour)

	_, err := pipe.Exec()
	fmt.Println(incr.Val(), err)
}
