package main

// Demo program that acts as a server.  It listens, connects, sends a single
// line message, and closes the connection.  (wash, rinse, repeat)
//

import (
	"crypto/tls"
	"flag"
	"log"
	_ "net"
)

const (
	PORT = ":8080"
)

func main() {

	certFile := flag.String("cert", "cert.pem", "certificate chain (PEM)")
	keyFile := flag.String("key", "key.pem", "certificate key (PEM)")
	flag.Parse()

	log.Printf("loading cert from:  %s\n", *certFile)
	log.Printf("loading key from:   %s\n", *keyFile)

	// load certificate + key from PEM

	cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
	checkErr("failed to load cert and key files", err)
	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}

	// listen & serve

	// l, err := net.Listen("tcp", PORT)   // non-TLS method

	l, err := tls.Listen("tcp", PORT, cfg) // TLS Listen
	checkErr("failed to listen", err)
	defer l.Close()
	log.Print("listening ...\n")

	conn, err := l.Accept()
	checkErr("err accepting on listener", err)
	log.Print("connected\n")

	n, err := conn.Write([]byte("response from server\n"))
	checkErr("error writing to connection", err)
	log.Printf("successfully wrote %d bytes", n)

	err = conn.Close()
	checkErr("error closing connection", err)
	log.Print("connection closed\n")
	log.Print("done.\n")
}

func checkErr(msg string, err error) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err.Error())
	}
}
