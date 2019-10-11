package main

// Example:  High Level 'reader' to Kafka
//
//  https://godoc.org/github.com/segmentio/kafka-go#Reader
//

import (
	"context"
	"fmt"
	"log"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

func main() {

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:         []string{"localhost:9092"},
		Topic:           "sandbox-topic",
		Partition:       0,
		MinBytes:        10e3,
		MaxBytes:        10e6,
		MaxWait:         time.Second,
		ReadLagInterval: -1,
	})
	defer r.Close()
	r.SetOffset(0)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Msg(offset: %d) - k: %s | v: %s\n", m.Offset, string(m.Key), string(m.Value))
	}
}
