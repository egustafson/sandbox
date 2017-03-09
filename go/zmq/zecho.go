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
	err = sock.Bind("tcp://*:5555")
	if err != nil {
		log.Fatal(err)
	}

	// Loop forever
	for {
		msg, err := sock.Recv(0)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Received: %s", msg)
		// just send the same message back
		_, err = sock.Send(msg, 0)
		if err != nil {
			log.Fatal(err)
		}
	}
}
