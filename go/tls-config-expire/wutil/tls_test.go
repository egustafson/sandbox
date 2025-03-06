package wutil_test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/egustafson/sandbox/go/tls-config-expire/wutil"
)

func TestAllNotBefore_nil(t *testing.T) {

	tlsProfile := &wutil.TlsProfile{DisableTls: true}
	tlsConfig, err := wutil.MakeTlsConfig(tlsProfile)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = tlsConfig.AllNotBefore()
	if err == nil {
		t.Fatal("expected error response")
	}
}

func TestAllNotAfter_nil(t *testing.T) {

	tlsProfile := &wutil.TlsProfile{DisableTls: true}
	tlsConfig, err := wutil.MakeTlsConfig(tlsProfile)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = tlsConfig.AllNotAfter()
	if err == nil {
		t.Fatal("expected error response")
	}
}

func TestMakeingCertificates(t *testing.T) {
	now := time.Now()

	caCert, caKey := makeCertAndKey(now.Add(-time.Hour), now.Add(time.Hour), nil, nil, 0)
	intCert, intKey := makeCertAndKey(now.Add(-time.Hour), now.Add(time.Hour), caCert, caKey, 1)
	leafCert, leafKey := makeCertAndKey(now.Add(-time.Hour), now.Add(time.Hour), intCert, intKey, 2)

	_ = leafCert
	_ = leafKey

	roots := x509.NewCertPool()
	roots.AddCert(caCert)

	opts := x509.VerifyOptions{Roots: roots}
	_, err := intCert.Verify(opts)
	if err != nil {
		t.Fatal(err.Error())
	}
}

//
// Certificate generation tools
//

func makeCertAndKey(start, end time.Time, signer *x509.Certificate, signerKey any, ser int64) (cert *x509.Certificate, key *rsa.PrivateKey) {

	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	template := x509.Certificate{
		SerialNumber:          big.NewInt(ser),
		Subject:               pkix.Name{Organization: []string{"Test Org"}, CommonName: fmt.Sprintf("cert-sn-%d", ser)},
		NotBefore:             start,
		NotAfter:              end,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
	}

	if signer == nil { // then self-sign the certificate
		signer = &template
		signerKey = key
	}

	certBytes, err := x509.CreateCertificate(
		rand.Reader,
		&template,
		signer,
		&key.PublicKey,
		signerKey,
	)
	if err != nil {
		panic(err)
	}

	cert, err = x509.ParseCertificate(certBytes)
	if err != nil {
		panic(err)
	}

	const debug = false
	if debug {
		certPEM := pem.EncodeToMemory(&pem.Block{
			Type:  "CERTIFICATE",
			Bytes: certBytes,
		})
		keyPem := pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		})
		fmt.Println(string(certPEM))
		fmt.Println(string(keyPem))
	}

	return cert, key
}
