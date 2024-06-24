package main

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func producer(n string) {
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), network, host, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	for {
		<-time.After(messageWriteTime)
		user, randomUserErr := getRandomUser()
		if randomUserErr == nil {
			buffer, randomUserErr := json.Marshal(
				Message{
					Channel:  n,
					UserInfo: user,
				})
			if randomUserErr == nil {
				_, err = conn.Write(buffer)
			} else {
				_, err = conn.Write([]byte("unhandled error"))
			}
		} else {
			_, err = conn.Write([]byte(randomUserErr.Error()))
		}
	}
}
