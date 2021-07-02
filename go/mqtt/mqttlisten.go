package main

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func onMessagePub(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("topic: %s\n  %s\n", msg.Topic(), msg.Payload())
}

func onConnect(client mqtt.Client) {
	fmt.Println("connected.")
}

func onDisconnect(client mqtt.Client, err error) {
	fmt.Printf("connection lost: %v\n", err)
}

const (
	broker    = "tcp://localhost:1883"
	client_id = "sandbox-test-client"
	topic     = "#"
	qos       = 1
)

func main() {
	options := mqtt.NewClientOptions()
	options.AddBroker(broker)
	options.SetClientID(client_id)
	options.SetDefaultPublishHandler(onMessagePub)
	options.OnConnect = onConnect
	options.OnConnectionLost = onDisconnect

	client := mqtt.NewClient(options)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	token = client.Subscribe(topic, qos, nil)
	token.Wait()
	fmt.Printf("subscribed to topic: %s\n", topic)

	time.Sleep(1 * time.Minute)
	client.Disconnect(100) // linger 100ms before closing
	fmt.Println("done.")
}
