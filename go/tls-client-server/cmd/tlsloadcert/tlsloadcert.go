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

const (
	CAFILE  = "demo-ca-cert.pem"
	SRVCERT = "demo-srv-cert.pem"
	SRVKEY  = "demo-srv-key.pem"
)

func main() {

	caFile := flag.String("ca", CAFILE, "certificate authority (PEM)")
	certFile := flag.String("cert", SRVCERT, "certificate chain (PEM)")
	keyFile := flag.String("key", SRVKEY, "certificate key (PEM)")
	flag.Parse()

	fmt.Printf("loading CA from:    %s\n", *caFile)
	fmt.Printf("loading cert from:  %s\n", *certFile)
	fmt.Printf("loading key from:   %s\n", *keyFile)

	// --  Load CA  ----------

	caPem, err := ioutil.ReadFile(*caFile)
	checkError("could not read CA file: ", err)

	caBlock, _ := pem.Decode(caPem)
	if caBlock == nil {
		log.Fatal("failed to decode the CA PEM block")
	}
	fmt.Printf("CA block is: %s\n", caBlock.Type)

	ca, err := x509.ParseCertificate(caBlock.Bytes)
	checkError("failed to parse CA: ", err)

	_ = ca // the CA (used when creating a listener)

	// --  Load Cert  ----------

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

	srvCert, err := x509.ParseCertificate(block.Bytes)
	checkError("failed to parse certificate: ", err)

	_ = srvCert // the server certificate (unused)

	// --  Load Key  ----------

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

	srvKey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		log.Fatal("could not cast parsedKey to *rsa.PrivateKey")
	}

	_ = srvKey // the server private key (unused)

	// ------------------------------------------------------
	// Load Cert + Key and allocate listener (alternate test)

	cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
	checkError("failed to load cert and key files: ", err)

	caCertPool := x509.NewCertPool()
	caCertPool.AddCert(ca) // loaded in the 'load CA' section above

	cfg := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,                     // used by clients to verify servers (not necessary here)
		ClientAuth:   tls.RequireAndVerifyClientCert, // the most restrictive
		ClientCAs:    caCertPool,                     // used by servers to verify clients
	}
	listener, err := tls.Listen("tcp", ":8080", cfg)

	_ = listener

	fmt.Print("ok.\n")
}

func checkError(msg string, err error) {
	if err != nil {
		log.Fatal(msg + err.Error())
	}
}
