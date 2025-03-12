package wutil

import (
	"bytes"
	"crypto"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"strings"
)

// This file defines holder functions for both an x.509 Certificate and a
// Private Key.

// Cert represents a single X.509 Certificate
type Cert []byte

func LoadCertFromDER(der []byte) (Cert, error) {
	if _, err := x509.ParseCertificate(der); err != nil {
		// decoded bytes did not parse as a CERTIFICATE
		return nil, err
	}
	return bytes.Clone(der), nil
}

func LoadCertFromPEM(pemData []byte) (Cert, error) {
	block, _ := pem.Decode(pemData)
	if block == nil || block.Type != "CERTIFICATE" {
		return nil, errors.New("data did not contain a certificate PEM block")
	}
	if _, err := x509.ParseCertificate(block.Bytes); err != nil {
		return nil, err
	}
	return block.Bytes, nil
}

func (der Cert) AsDER() []byte {
	return []byte(der)
}

func (der Cert) AsPEM() []byte {
	block := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: []byte(der),
	}
	return pem.EncodeToMemory(block)
}

func (der Cert) AsX509Certificate() *x509.Certificate {
	cert, _ := x509.ParseCertificate(der)
	return cert
}

// Key represents a private Key used as part of X.509 PKI
type Key struct {
	k crypto.PrivateKey
}

func LoadKeyFromDER(der []byte) (*Key, error) {
	if k, err := x509.ParsePKCS8PrivateKey(der); err == nil {
		return &Key{k: k}, nil
	} else {
		return nil, err
	}
}

func LoadKeyFromPEM(pemData []byte) (*Key, error) {
	block, _ := pem.Decode(pemData)
	if block == nil || !strings.Contains(block.Type, "PRIVATE KEY") {
		return nil, errors.New("data did not contain a private key PEM block")
	}
	return LoadKeyFromDER(block.Bytes)
}

func (key *Key) AsDER() []byte {
	der, _ := x509.MarshalPKCS8PrivateKey(key.k)
	return der
}

func (key *Key) AsPEM() []byte {
	der, _ := x509.MarshalPKCS8PrivateKey(key.k)
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: der,
	}
	return pem.EncodeToMemory(block)
}

func (key *Key) AsPrivateKey() crypto.PrivateKey {
	return &key.k
}
