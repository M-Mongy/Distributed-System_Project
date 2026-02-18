package main

import (
	"context"
	"log"
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

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
   // Prometheus metric for processed messages
   var messagesProcessed = prometheus.NewCounter(
	   prometheus.CounterOpts{
		   Name: "consumer_messages_processed_total",
		   Help: "Total number of messages processed by the consumer.",
	   },
   )
   prometheus.MustRegister(messagesProcessed)

   // Start metrics server in a goroutine
   go func() {
	   http.Handle("/metrics", promhttp.Handler())
	   log.Println("Consumer metrics endpoint on :2113/metrics")
	   http.ListenAndServe(":2113", nil)
   }()

   log.Println("Consumer started...")

   for {
	   msg, err := reader.ReadMessage(context.Background())
	   if err != nil {
		   log.Println("Error reading message:", err)
		   continue
	   }
	   log.Printf("Received: %s = %s\n", string(msg.Key), string(msg.Value))
	   messagesProcessed.Inc()
   }
}
