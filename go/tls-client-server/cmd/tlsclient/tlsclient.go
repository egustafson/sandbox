package main

// Demo program that dials a server, reads the first line sent by the server,
// prints that line to stdout, closes and exits.
//
// Protocol followed:
//  1.  Connect
//  2.  read to EOL
//  2a.  - print what is read
//  3.  Close connection
//  4.  exit

import (
	"bufio"
	"crypto/tls"
	"flag"
	"log"
	_ "net"
)

const (
	EP = "localhost:8080"
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

	// dial and recv message

	// conn, err := net.Dial("tcp", EP)    // non-TLS method

	conn, err := tls.Dial("tcp", EP, cfg)
	checkErr("Dial failed", err)

	rd := bufio.NewReader(conn)
	msg, err := rd.ReadString('\n')
	checkErr("failed reading", err)

	err = conn.Close()
	checkErr("error on close", err)

	log.Printf("msg recv: %s", msg)
	log.Print("done.")
}

func checkErr(msg string, err error) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err.Error())
	}
}
