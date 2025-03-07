package wutil_test

import (
	"crypto/x509"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/egustafson/sandbox/go/tls-config-expire/wutil"
)

func TestMakeCertAndKey(t *testing.T) {

	caCert, caKey := wutil.MakeCertAndKey() // <-- self-signed
	assert.NotNil(t, caCert)
	assert.NotNil(t, caKey)

	cert, key := wutil.MakeCertAndKey(wutil.WithSigner(caCert, caKey))
	assert.NotNil(t, cert)
	assert.NotNil(t, key)

	roots := x509.NewCertPool()
	roots.AddCert(caCert)
	opts := x509.VerifyOptions{
		Roots: roots,
	}

	_, err := cert.Verify(opts)
	assert.Nil(t, err)
}

func TestMakeCertChain(t *testing.T) {

	chain := wutil.MakeCertChain()
	assert.NotNil(t, chain)

	assert.True(t, len(chain) == 2) // default length is 2

	// TODO: extract ca and leaf and validate leaf against ca just like in TestMakeCertAndKey
}

// TODO: write a test that creates a longer chain (5 ish elements) and validate the full chain
