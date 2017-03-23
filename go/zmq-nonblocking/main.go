package main

import (
	zmq "github.com/pebbe/zmq4"
	"log"
	"syscall"
)

func main() {

	log.Print("Starting example.")
	sock, _ := zmq.NewSocket(zmq.PAIR)
	sock.Bind("inproc://any-addr")

	msg, err := sock.RecvMessageBytes(zmq.DONTWAIT) // this should err
	if err != nil {
		if zmq.AsErrno(err) == zmq.Errno(syscall.EAGAIN) {
			log.Print("* non-blocking call:  nothing to read. *")
		}
	}

	// and for edification
	log.Printf("msg: %v", msg)
	log.Printf("error:  %v", err)
	log.Printf("errno:  %d", zmq.AsErrno(err))
	log.Printf("EAGAIN: %d", zmq.Errno(syscall.EAGAIN))
}
