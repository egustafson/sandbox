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
	"crypto/x509"
	"encoding/pem"
	"flag"
	"io/ioutil"
	"log"
	_ "net"
)

const (
	EP = "localhost:8080"

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

	// --  dial & receive  ---------------------------------

	secureCfg := &tls.Config{
		Certificates: []tls.Certificate{cert}, // cert(and key) to use for TLS
		RootCAs:      caCertPool,              // used by client to verify server, client verifies by default
	}

	noCfg := &tls.Config{
		InsecureSkipVerify: true, // no checks, self generate unsigned cert/key
	}
	_ = noCfg // not used (but could be, in the tlsDial() call)

	// conn, err := net.Dial("tcp", EP)    // non-TLS method
	conn, err := tls.Dial("tcp", EP, secureCfg) // TLS Dial
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
