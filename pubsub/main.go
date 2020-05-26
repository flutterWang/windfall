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

	pubsub := client.Subscribe("mychannel1")

	ch := pubsub.Channel()

	err := client.Publish("mychannel1", "hello world").Err()
	if err != nil {
		panic(err)
	}

	time.AfterFunc(time.Second, func() {
		// When pubsub is closed channel is closed too.
		_ = pubsub.Close()
	})

	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}
}
