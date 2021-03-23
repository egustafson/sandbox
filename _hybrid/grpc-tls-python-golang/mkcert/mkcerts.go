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
	cafile   = "test-ca.pem"
	certfile = "test-cert.pem"
	keyfile  = "test-key.pem"

	filemode = 0644 // u+rw,go+r - output mode of PEM files
	lifetime = 10   // days - duration of all created keys/certs
)

func main() {
	ca, caKey, caBytes := makeCA()
	certBytes, keyBytes := mkCert(ca, caKey)

	_ = caKey // the CA key is DISCARDED

	writePem(cafile, "CERTIFICATE", caBytes)
	writePem(certfile, "CERTIFICATE", certBytes)
	writePem(keyfile, "RSA PRIVATE KEY", keyBytes)
}

// Create a CA with signing (private) key.
func makeCA() (ca *x509.Certificate, caKey *rsa.PrivateKey, caBytes []byte) {
	var err error

	ca = &x509.Certificate{
		SerialNumber: big.NewInt(1000),
		Subject: pkix.Name{
			Organization:  []string{"Bogus Example CA"},
			Country:       []string{"US"},
			Locality:      []string{"Denver"},
			StreetAddress: []string{"1 Any Street"},
			PostalCode:    []string{"80001"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(0, 0, lifetime),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}
	caKey, err = rsa.GenerateKey(rand.Reader, 4096)
	checkErr("failed generating CA key", err)

	caBytes, err = x509.CreateCertificate(rand.Reader, ca, ca, &caKey.PublicKey, caKey)
	checkErr("failed to create (self sign) CA certificate", err)

	return ca, caKey, caBytes
}

// Create a client or server certificate and private key, signed by the CA
func mkCert(ca *x509.Certificate, caKey *rsa.PrivateKey) (certBytes []byte, keyBytes []byte) {
	cert := &x509.Certificate{
		SerialNumber: big.NewInt(1000),
		Subject: pkix.Name{
			Organization:  []string{"Bogus Example Org"},
			Country:       []string{"US"},
			Locality:      []string{"Denver"},
			StreetAddress: []string{"1 Any Street"},
			PostalCode:    []string{"80001"},
		},
		DNSNames:     []string{"localhost", "localhost.local"},
		IPAddresses:  []net.IP{net.IPv4(127, 0, 0, 1), net.IPv6loopback},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(0, 0, lifetime),
		SubjectKeyId: []byte{1, 2, 3, 4, 9},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}
	certPrivKey, err := rsa.GenerateKey(rand.Reader, 4096)
	checkErr("failed generating srv cert key", err)

	certBytes, err = x509.CreateCertificate(rand.Reader, cert, ca, &certPrivKey.PublicKey, caKey)
	checkErr("failed to sign the srv certificate with CA", err)

	keyBytes = x509.MarshalPKCS1PrivateKey(certPrivKey)

	return certBytes, keyBytes
}

func writePem(fn string, pemType string, pemBytes []byte) {
	blk := &pem.Block{
		Type:  pemType,
		Bytes: pemBytes,
	}
	pemBuf := new(bytes.Buffer)
	err := pem.Encode(pemBuf, blk)
	checkErr("problem encoding PEM", err)
	err = ioutil.WriteFile(fn, pemBuf.Bytes(), filemode)
	checkErr("problemwriting PEM file", err)
}

// Error handler
func checkErr(msg string, err error) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err.Error())
	}
}
