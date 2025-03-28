package wutil_test

import (
	"crypto/x509"
	"fmt"
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

	var chainLen []uint = []uint{1, 2, 3, 4, 5, 6}
	for _, l := range chainLen {
		name := fmt.Sprintf("chain-len-%d", l)
		t.Run(name, func(t *testing.T) {

			chain := wutil.MakeCertChain(wutil.WithLength(l))
			assert.NotNil(t, chain)

			assert.True(t, uint(len(chain)) == l) // default length is 2

			roots := x509.NewCertPool()
			roots.AddCert(chain.CA().Cert)
			intermediates := x509.NewCertPool()
			ii := chain.Intermediates()
			for _, i := range ii {
				intermediates.AddCert(i.Cert)
			}
			opts := x509.VerifyOptions{
				Roots:         roots,
				Intermediates: intermediates,
			}

			leaf := chain.Leaf().Cert
			_, err := leaf.Verify(opts)
			assert.Nil(t, err)
		})
	}
}
