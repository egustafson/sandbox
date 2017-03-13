//
// Demonstrate setting the identity of a connecting DEALER (or REQ) socket
// when connecting through a ROUTER.   This method could be used to explicitly
// de-mux messages from multiple, connected ZMQ sockets.  With DEALER as the
// sender, messages do not need to follow the strict REQ-REP pattern.
//
// This is an almost direct copy from the zguide (zguide.zeromq.org/go:identity)
//

package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"log"
)

var (
	bindaddr = "inproc://example"
)

func dump(sink *zmq.Socket) {
	parts, err := sink.RecvMessage(0)
	if err != nil {
		log.Print(err)
	}
	fmt.Println("--")
	for _, msgdata := range parts {
		is_text := true
		fmt.Printf("[%03d] ", len(msgdata))
		for _, char := range msgdata {
			if char < 32 || char > 127 {
				is_text = false
			}
		}
		if is_text {
			fmt.Printf("%s\n", msgdata)
		} else {
			fmt.Printf("%X\n", msgdata)
		}
	}
}

func sendto(sock *zmq.Socket, msg string) {
	c, err := sock.Send(msg, 0)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Sent %d bytes.", c)
}

func main() {
	sink, err := zmq.NewSocket(zmq.ROUTER)
	if err != nil {
		log.Fatal(err)
	}
	defer sink.Close()
	sink.Bind(bindaddr)

	// First allow ZMQ to set the socket identity in the ROUTER
	anonsock, err := zmq.NewSocket(zmq.DEALER)
	defer anonsock.Close()
	if err != nil {
		log.Fatal(err)
	}
	anonsock.Connect(bindaddr)

	// acutally send
	sendto(anonsock, "ROUTER uses a generated 5 byte identity")
	dump(sink)

	// Then set the identity ourselves
	idsock, err := zmq.NewSocket(zmq.DEALER)
	if err != nil {
		log.Fatal(err)
	}
	defer idsock.Close()
	idsock.SetIdentity("IDSOCK-1")
	idsock.Connect(bindaddr)

	// send with an explicitly identified socket
	sendto(idsock, "ROUTER socket uses DEALER socket's identity")
	//
	// note - with DEALER we are not confined to req/rep sequencing.
	sendto(idsock, "msg-2 to idsock")
	sendto(idsock, "msg-3 to idsock")
	dump(sink)
	dump(sink)
	dump(sink)

	// done.
	fmt.Println("done.")
}
