package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	sarama "github.com/Shopify/sarama"
)

const (
	kafkaConn = "127.0.0.1:9092"
	topic     = "senz"
)

func main() {
	producer, err := initProducer()
	if err != nil {
		fmt.Println("Error producer: ", err.Error())
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter msg: ")
		msg, _ := reader.ReadString('\n')

		publish(msg, producer)

		// (alternative) publish with go routine
		// go publish(msg, producer)
	}
}

func initProducer() (sarama.SyncProducer, error) {
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	// (alternative) async producer
	// prd, err := sarama.NewAsyncProducer([]string{kafkaConn}, config)

	// sync producer
	prd, err := sarama.NewSyncProducer([]string{kafkaConn}, config)

	return prd, err
}

func publish(message string, producer sarama.SyncProducer) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	p, o, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("Error publish: ", err.Error())
	}

	// (alternative) publish async
	// producer.Input() <- &sarama.ProducerMessage{ ...

	fmt.Println("Partition: ", p)
	fmt.Println("Offset: ", o)
}
