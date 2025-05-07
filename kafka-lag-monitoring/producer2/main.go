package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Kafka writer (producer)
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "topic2",
	})

	// Kafka reader (consumer)
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "topic2",
		GroupID: "my-app-group",
	})

	// Produce messages
	go func() {
		for {
			err := writer.WriteMessages(context.Background(),
				kafka.Message{
					Key:   []byte("key2"),
					Value: []byte("Bye message " + time.Now().String()),
				},
			)
			if err != nil {
				log.Println("Error writing message:", err)
			}
			time.Sleep(2 * time.Second)
		}
	}()

	// Consume messages with delay to simulate lag
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}
		fmt.Printf("Message: %s\n", string(m.Value))
		time.Sleep(5 * time.Second) // Delay to simulate lag
	}
}
