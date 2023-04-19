package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	topic          = "myTopic"
	broker1Address = "localhost:9092"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go consumer()
	fmt.Println("NOT LOCK THREAD")
	wg.Wait()
}

func consumer() {
	fmt.Println("CONSUMER INITIALIZED")
	ctx := context.Background()
	consumer := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{broker1Address},
		Topic:          topic,
		GroupID:        "my-group",
		CommitInterval: time.Duration(time.Second * 60),
	})
	defer consumer.Close()
	for {
		msg, err := consumer.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		fmt.Println("received: ", string(msg.Value))
	}
}
