package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {

	//Consumer

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		GroupID:  "consumer",
		Topic:    "quickstart",
		MinBytes: 0,
		MaxBytes: 10e6,
	})
	defer reader.Close()
	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			reader.Close()
			log.Fatal("cannot reaceive a message")
		}

		fmt.Println("receive a message: ", string(message.Value))
	}
}
