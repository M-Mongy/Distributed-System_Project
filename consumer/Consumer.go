package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

var reader *kafka.Reader

func init() {
	log.Println("Initializing Kafka consumer...")

	reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "user_events",
		GroupID: "consumer-group-1",
	})
}

func main() {
	log.Println("Consumer started...")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}
		log.Printf("Received: %s = %s\n", string(msg.Key), string(msg.Value))
	}
}
