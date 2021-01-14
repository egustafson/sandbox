package main

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	certFile := flag.String("cert", "cert.pem", "certificate chain (PEM)")
	keyFile := flag.String("key", "key.pem", "certificate key (PEM)")
	flag.Parse()

	fmt.Printf("loading cert from:  %s\n", *certFile)
	fmt.Printf("loading key from:   %s\n", *keyFile)

	// Cert

	certPem, err := ioutil.ReadFile(*certFile)
	checkError("could not read certFile: ", err)

	block, rest := pem.Decode(certPem)
	if block == nil {
		log.Fatal("failed to decode 1st PEM block")
	}
	fmt.Printf("initial block is:  %s\n", block.Type)
	if rest == nil {
		log.Fatal("did not decode any additional keys in the cert")
	}

	_, err = x509.ParseCertificate(block.Bytes)
	checkError("failed to parse certificate: ", err)

	// Key

	keyPem, err := ioutil.ReadFile(*keyFile)
	checkError("could not read keyFile: ", err)

	kBlock, _ := pem.Decode(keyPem)
	if block == nil {
		log.Fatal("failed to decode key PEM block")
	}
	fmt.Printf("key block is:  %s\n", kBlock.Type)

	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS1PrivateKey(kBlock.Bytes); err != nil {
		log.Print("key is not a PKCS1 format key")
		if parsedKey, err = x509.ParsePKCS8PrivateKey(kBlock.Bytes); err != nil {
			log.Fatal("key is not a PKCS8 format key - abandoing.")
		}
	}

	_, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		log.Fatal("could not cast parsedKey to *rsa.PrivateKey")
	}

	// Load Cert + Key and allocate listener (alternate test)

	cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
	checkError("failed to load cert and key files: ", err)

	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}
	listener, err := tls.Listen("tcp", ":8080", cfg)

	_ = listener

	fmt.Print("ok.\n")
}

func checkError(msg string, err error) {
	if err != nil {
		log.Fatal(msg + err.Error())
	}
}
