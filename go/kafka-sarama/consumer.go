package main

import (
	"fmt"
	"log"
	"os"
	"time"

	sarama "github.com/Shopify/sarama"
	consumergroup "github.com/wvanbergen/kafka/consumergroup"
)

const (
	zookeeperConn = "127.0.0.1:2181"
	cgroup        = "zgroup"
	topic         = "senz"
)

func main() {
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	cg, err := initConsumer()
	if err != nil {
		fmt.Println("Error consumer group: ", err.Error())
		os.Exit(1)
	}
	defer cg.Close()

	// run consumer
	consume(cg)
}

func initConsumer() (*consumergroup.ConsumerGroup, error) {
	config := consumergroup.NewConfig()
	config.Offsets.Initial = sarama.OffsetOldest
	config.Offsets.ProcessingTimeout = 10 * time.Second

	cg, err := consumergroup.JoinConsumerGroup(cgroup, []string{topic}, []string{zookeeperConn}, config)
	if err != nil {
		return nil, err
	}
	return cg, err
}

func consume(cg *consumergroup.ConsumerGroup) {
	for {
		select {
		case msg := <-cg.Messages():
			if msg.Topic != topic {
				continue
			}
			fmt.Println("Topic: ", msg.Topic)
			fmt.Println("Value: ", string(msg.Value))

			err := cg.CommitUpto(msg)
			if err != nil {
				fmt.Println("Error commit zookeeper: ", err.Error())
			}
		}
	}
}
