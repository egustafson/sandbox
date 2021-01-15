package main

// Demo program that acts as a server.  It listens, connects, sends a single
// line message, and closes the connection.  (wash, rinse, repeat)
//

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"io/ioutil"
	"log"
	_ "net"
)

const (
	PORT = ":8080"

	CAFILE  = "demo-ca-cert.pem"
	SRVCERT = "demo-srv-cert.pem"
	SRVKEY  = "demo-srv-key.pem"
)

func main() {

	caFile := flag.String("ca", CAFILE, "certificate authority (PEM)")
	certFile := flag.String("cert", SRVCERT, "certificate chain (PEM)")
	keyFile := flag.String("key", SRVKEY, "certificate key (PEM)")
	flag.Parse()

	log.Printf("loading CA from:    %s\n", *caFile)
	log.Printf("loading cert from:  %s\n", *certFile)
	log.Printf("loading key from:   %s\n", *keyFile)

	// --  load certificate + key from PEM  ----------------

	cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
	checkErr("failed to load cert and key files", err)

	// --  load the CA from PEM file -----------------------

	caPem, err := ioutil.ReadFile(*caFile)
	checkErr("failed to read CA file", err)

	caBlock, _ := pem.Decode(caPem)
	if caBlock == nil {
		log.Fatal("failed to decode the CA PEM block")
	}

	ca, err := x509.ParseCertificate(caBlock.Bytes)
	checkErr("failed to parse CA from PEM block", err)
	caCertPool := x509.NewCertPool()
	caCertPool.AddCert(ca)

	// --  listen & serve  ---------------------------------

	secureCfg := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert, // the most restrictive (verify against ClientCAs)
		ClientCAs:    caCertPool,                     // used by servers to verify clients
	}

	openCfg := &tls.Config{
		InsecureSkipVerify: true, // disables all checks -- opposite of tls.RequireAndVerify...
	}
	_ = openCfg // not used (but could be, in the tls.Listen() call)

	// l, err := net.Listen("tcp", PORT)   // non-TLS method
	l, err := tls.Listen("tcp", PORT, secureCfg) // TLS Listen
	checkErr("failed to listen", err)
	defer l.Close()
	log.Print("listening ...\n")

	// --  accept & process (and close)  -------------------

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
