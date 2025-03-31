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

func TestTlsProfile_NotBefore_DisableTls(t *testing.T) {

	tlsProfile := &wutil.TlsProfile{DisableTls: true}
	_, err := tlsProfile.TlsConfig()
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = tlsProfile.NotBefore()
	if err == nil {
		t.Fatal("expected error response")
	}
}

func TestTlsProfile_NotAfter_DisableTls(t *testing.T) {

	tlsProfile := &wutil.TlsProfile{DisableTls: true}
	_, err := tlsProfile.TlsConfig()
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = tlsProfile.NotAfter()
	if err == nil {
		t.Fatal("expected error response")
	}
}

func TestMakingCertificates(t *testing.T) {
	now := time.Now()

	caCert, caKey := makeCertAndKey(now.Add(-time.Hour), now.Add(time.Minute), nil, nil, 0)
	intCert, intKey := makeCertAndKey(now.Add(-time.Hour), now.Add(time.Minute), caCert, caKey, 1)
	leafCert, leafKey := makeCertAndKey(now.Add(-time.Hour), now.Add(time.Minute), intCert, intKey, 2)

	_ = leafCert
	_ = leafKey

	intermediates := x509.NewCertPool()
	intermediates.AddCert(intCert)

	roots := x509.NewCertPool()
	roots.AddCert(caCert)

	opts := x509.VerifyOptions{
		Roots:         roots,
		Intermediates: intermediates,
	}

	_, err := leafCert.Verify(opts)
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
		IsCA:                  true,
		DNSNames:              []string{"localhost"},
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
