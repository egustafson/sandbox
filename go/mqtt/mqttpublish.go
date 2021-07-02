package main

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func onConnect(client mqtt.Client) {
	fmt.Println("connected.")
}

func onDisconnect(client mqtt.Client, err error) {
	fmt.Printf("disconnected: %v\n", err)
}

const (
	broker    = "tcp://localhost:1883"
	client_id = "sandbox-test-publisher"
	topic     = "sandbox/test-publish"
	qos       = 0
	retension = false
)

func main() {
	options := mqtt.NewClientOptions()
	options.AddBroker(broker)
	options.SetClientID(client_id)
	options.OnConnect = onConnect
	options.OnConnectionLost = onDisconnect

	client := mqtt.NewClient(options)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	payload := "sandbox-test-payload"
	token = client.Publish(topic, qos, retension, payload)
	token.Wait()

	client.Disconnect(100) // linger 100ms
	fmt.Println("done.")
}
