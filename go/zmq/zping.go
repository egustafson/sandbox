package main

import (
	zmq "github.com/pebbe/zmq4"
	"log"
)

func main() {
	ctx, err := zmq.NewContext()
	if err != nil {
		log.Fatal(err)
	}
	defer ctx.Term()
	sock, err := ctx.NewSocket(zmq.PAIR)
	if err != nil {
		log.Fatal(err)
	}
	defer sock.Close()
	err = sock.Connect("tcp://localhost:5555")
	if err != nil {
		log.Fatal(err)
	}

	// send message read reply
	msg := "Ping"
	_, err = sock.Send(msg, 0)
	if err != nil {
		log.Fatal(nil)
	}
	log.Printf("Sent: '%s'", msg)
	msg, err = sock.Recv(0)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Recv: '%s'", msg)
}
