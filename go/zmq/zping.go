package main

import (
	zmq "github.com/pebbe/zmq4"
	"log"
)

var (
	srvAddr = "tcp://localhost:5555"
)

func main() {

	// create a new socket for communicating.
	sock, err := zmq.NewSocket(zmq.PAIR)
	if err != nil {
		log.Fatal(err)
	}
	defer sock.Close()

	// connect to the "server"
	log.Printf("connecting to:  %s", srvAddr)
	err = sock.Connect(srvAddr)
	if err != nil {
		log.Fatal(err)
	}

	// send message then read reply
	msg := "Ping"
	log.Printf("sending msg: %s", msg)
	_, err = sock.Send(msg, 0)
	if err != nil {
		log.Fatal(nil)
	}
	log.Printf("Sent: '%s', receiving ...", msg)
	msg, err = sock.Recv(0)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Recv: '%s'", msg)
}
