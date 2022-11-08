package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nsqio/go-nsq"
	"github.com/spf13/cobra"
)

var subCmd = &cobra.Command{
	Use:   "sub",
	Short: "Subscribe to messages",
	Run:   doSub,
}

func init() {
	rootCmd.AddCommand(subCmd)
}

type demoMsgHandler struct{}

func (h *demoMsgHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		return nil // signals message has been processed
	}

	fmt.Printf(": %s\n", string(m.Body))
	return nil
}

func doSub(cmd *cobra.Command, args []string) {
	fmt.Println("subscribing")

	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("topic-1", "channel-a", config)
	if err != nil {
		log.Fatal(err)
	}

	consumer.AddHandler(&demoMsgHandler{})

	err = consumer.ConnectToNSQLookupd("dev0.elfwerks:4161")
	if err != nil {
		log.Fatal(err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan // block on signal

	consumer.Stop()
}
