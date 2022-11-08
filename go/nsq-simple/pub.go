package main

import (
	"fmt"
	"log"

	"github.com/nsqio/go-nsq"
	"github.com/spf13/cobra"
)

var pubCmd = &cobra.Command{
	Use:   "pub",
	Short: "Publish a single message",
	Run:   doPub,
}

func init() {
	rootCmd.AddCommand(pubCmd)
}

func doPub(cmd *cobra.Command, args []string) {
	fmt.Println("publish")

	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("dev0.elfwerks:4150", config)
	if err != nil {
		log.Fatal(err)
	}

	msgBody := []byte("message-1")
	topicName := "topic-1"

	err = producer.Publish(topicName, msgBody)
	if err != nil {
		log.Fatal(err)
	}

	producer.Stop()
}
