package main

import (
	"context"

	"github.com/segmentio/kafka-go"
)

func main() {
	writer := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "quickstart",
	}

	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Value: []byte("Hello world from kafka1"),
		},
		kafka.Message{
			Value: []byte("Hello world from kafka2"),
		},
		kafka.Message{
			Value: []byte("Hello world from kafka3"),
		},
		kafka.Message{
			Value: []byte("Hello world from kafka4"),
		},
		kafka.Message{
			Value: []byte("Hello world from kafka5"),
		},
		kafka.Message{
			Value: []byte("Hello world from kafka6"),
		},
	)

	if err != nil {
		panic(err)
	}

}
