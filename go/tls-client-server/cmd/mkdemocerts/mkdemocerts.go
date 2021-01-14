package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"io/ioutil"
	"log"
	"math/big"
	"net"
	"time"
)

const (
	CAFILE    = "demo-ca-cert.pem"
	CAKEYFILE = "demo-ca-key.pem"
	SRVCERT   = "demo-srv-cert.pem"
	SRVKEY    = "demo-srv-key.pem"
)

func main() {
	ca, caKey := makeCA()
	srvCert, srvKey := makeSrvCert(ca, caKey)

	_ = srvCert
	_ = srvKey
}

// Create a Certificate Authority (CA) with signing (private) key.
func makeCA() (ca *x509.Certificate, caKey *rsa.PrivateKey) {
	ca = &x509.Certificate{
		SerialNumber: big.NewInt(1000),
		Subject: pkix.Name{
			Organization:  []string{"Test Company, INC."},
			Country:       []string{"US"},
			Locality:      []string{"Denver"},
			StreetAddress: []string{"123 Any Street"},
			PostalCode:    []string{"80000"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(0, 0, 2),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}
	caPrivKey, err := rsa.GenerateKey(rand.Reader, 4096)
	checkErr("failed generating CA key", err)

	caBytes, err := x509.CreateCertificate(rand.Reader, ca, ca, &caPrivKey.PublicKey, caPrivKey)
	checkErr("failed to create (self sign) CA certificate", err)

	// encode as PEM and write to file the CA and CA's key
	//
	writePem(CAFILE, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caBytes,
	})
	log.Printf("successfully wrote CA to '%s'", CAFILE)

	writePem(CAKEYFILE, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(caPrivKey),
	})
	log.Printf("successfully wrote CA key to '%s'", CAKEYFILE)

	// return the machine readable (PEM _decoded) CA + CA-Key
	return ca, caPrivKey
}

// Create a server certificate, and private key, signed by the CA.
func makeSrvCert(ca *x509.Certificate, caKey *rsa.PrivateKey) (srvCert *x509.Certificate, srvKey *rsa.PrivateKey) {
	cert := &x509.Certificate{
		SerialNumber: big.NewInt(2000),
		Subject: pkix.Name{
			Organization:  []string{"Test Company, INC."},
			Country:       []string{"US"},
			Locality:      []string{"Denver"},
			StreetAddress: []string{"123 Any Street"},
			PostalCode:    []string{"80000"},
		},
		IPAddresses:  []net.IP{net.IPv4(127, 0, 0, 1), net.IPv6loopback},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(0, 0, 2),
		SubjectKeyId: []byte{1, 2, 3, 4, 9},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}
	certPrivKey, err := rsa.GenerateKey(rand.Reader, 4096)
	checkErr("failed generating srv cert key", err)

	certBytes, err := x509.CreateCertificate(rand.Reader, cert, ca, &certPrivKey.PublicKey, caKey)
	checkErr("failed to sign the srv certificate with CA", err)

	// encode as PEM and write to file the server certificate and private key
	//
	writePem(SRVCERT, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})
	log.Printf("successfully wrote server cert to %s", SRVCERT)

	writePem(SRVKEY, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(certPrivKey),
	})
	log.Printf("successfully wrote server key to %s", SRVKEY)

	// return the machine readable (PEM _decoded_) server certificate and private key
	return cert, certPrivKey
}

func writePem(fn string, pb *pem.Block) {
	pemBuf := new(bytes.Buffer)
	err := pem.Encode(pemBuf, pb)
	checkErr("problem encoding PEM", err)
	err = ioutil.WriteFile(fn, pemBuf.Bytes(), 0644)
	checkErr("problem writing PEM file", err)
}

func checkErr(msg string, err error) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err.Error())
	}
}
