package main

import (
	zmq "github.com/pebbe/zmq4"
	"log"
)

var (
	monitor  = true
	monAddr  = "inproc://monitor.sock"
	bindAddr = "tcp://*:5555"
)

func sockMonitor(addr string) {
	msock, err := zmq.NewSocket(zmq.PAIR)
	if err != nil {
		log.Fatal(err)
	}
	defer msock.Close()
	defer log.Print("sockMonitor() - exited")
	err = msock.Connect(addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("-Tracing socket events")
	for {
		event, addr, val, err := msock.RecvEvent(0)
		if err != nil {
			log.Print(err)
			break
		}
		log.Printf("- event(%s) @%s [val: %d]", event, addr, val)
	}
}

func main() {

	// create the primary socket for communicating.
	sock, err := zmq.NewSocket(zmq.PAIR)
	if err != nil {
		log.Fatal(err)
	}
	defer sock.Close()

	// configure socket "monitoring"
	if monitor {
		err = sock.Monitor(monAddr, zmq.EVENT_ALL)
		if err != nil {
			log.Fatal(err)
		}
		go sockMonitor(monAddr)
	}

	// bind and accept connection(s)
	log.Printf("binding to: %s", bindAddr)
	err = sock.Bind(bindAddr)
	if err != nil {
		log.Fatal(err)
	}

	// Loop forever
	log.Print("serving echo() forever ...")
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
