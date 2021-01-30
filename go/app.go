package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/csarnataro/miniproject/messages"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// IncomingRedisChannel name of the redis channel
const IncomingRedisChannel = "postback-channel"

func main() {
	fmt.Println("listening to channel..." + IncomingRedisChannel)
	listenPubSubChannel()

}

func listenPubSubChannel() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer rdb.Close()

	pubsub := rdb.Subscribe(ctx, IncomingRedisChannel)

	subscriptionChannel := pubsub.Channel()
	defer pubsub.Close()

	for msg := range subscriptionChannel {

		var message messages.Message
		if err := json.Unmarshal([]byte(msg.Payload), &message); err != nil {
			log.Printf("Error on unmarshal JSON message %s", err)
			// should log wrong message received
		} else {
			go messages.ProcessMessage(message)
		}
	}
}
