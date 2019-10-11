package main

// Example:  High Level 'Writer' to Kafka
//
//  https://godoc.org/github.com/segmentio/kafka-go#Writer
//

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

func main() {

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      []string{"localhost:9092"},
		Topic:        "sandbox-topic",
		Balancer:     &kafka.LeastBytes{},
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		BatchSize:    1000,
		BatchTimeout: 5 * time.Millisecond,
	})
	defer w.Close()

	t := time.Now()
	for ii := 0; ii < 1000003; ii++ {
		key := fmt.Sprintf("%d%s", ii, t)
		khash := fmt.Sprintf("%x", sha1.Sum([]byte(key)))
		w.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(khash),
				Value: []byte("value-a"),
			},
		)
	}
}
