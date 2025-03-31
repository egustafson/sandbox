package wutil_test

import (
	"crypto/ecdsa"
	"strings"
	"testing"

	"github.com/egustafson/sandbox/go/tls-config-expire/wutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testCertPEM = `
-----BEGIN CERTIFICATE-----
MIIBqzCCAVKgAwIBAgIIKm0vI6YpsxMwCgYIKoZIzj0EAwIwJjEkMCIGA1UEAxMb
Y2VydC1zbi0zMDU3MTUxNTUyMjEyNTQ2MzIzMB4XDTI1MDMxMjAyNDEwNVoXDTI1
MDMxMjA0NDEwNVowJjEkMCIGA1UEAxMbY2VydC1zbi0zMDU3MTUxNTUyMjEyNTQ2
MzIzMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEs9mP1IWWbvojZzMIDoHoSERu
9TAIPOwCyJdANQmsIBAz8xGYaBxsYhgtqxB+67qIjjPQuELoZiUy655a2YnY/aNq
MGgwDwYDVR0PAQH/BAUDAwf/gDAPBgNVHSUECDAGBgRVHSUAMA8GA1UdEwEB/wQF
MAMBAf8wHQYDVR0OBBYEFBC1C8HiMZswtgOOttZlTUMF5sYJMBQGA1UdEQQNMAuC
CWxvY2FsaG9zdDAKBggqhkjOPQQDAgNHADBEAiAkDURr423XYs7WOWjGipPvTpjg
sNcABjQwvUFOp9XS8gIgFXdzYd615C6N+L1wQCBwwlivApWnfmzkFRKAW33UA54=
-----END CERTIFICATE-----
`
	testKeyPEM = `
-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIObZpf1QC2odHz335VDO+VApoNhVW91VCucB3JG20FpwoAoGCCqGSM49
AwEHoUQDQgAEs9mP1IWWbvojZzMIDoHoSERu9TAIPOwCyJdANQmsIBAz8xGYaBxs
YhgtqxB+67qIjjPQuELoZiUy655a2YnY/Q==
-----END EC PRIVATE KEY-----
`
)

func TestLoadCert(t *testing.T) {

	cert, err := wutil.LoadCertFromPEM([]byte(testCertPEM))
	assert.Nil(t, err)
	assert.True(t, len(cert) > 0)

	x509Cert := cert.AsX509Certificate()
	if assert.NotNil(t, x509Cert) {
		assert.Equal(t, "cert-sn-3057151552212546323", x509Cert.Subject.CommonName)
	}

	cert2, err := wutil.LoadCertFromDER(cert.AsDER())
	assert.Nil(t, err)
	assert.True(t, len(cert2) == len(cert))

	pem := cert.AsPEM()
	assert.True(t, strings.Contains(testCertPEM, string(pem)))
}

func TestLoadKey(t *testing.T) {

	key, err := wutil.LoadKeyFromPEM([]byte(testKeyPEM))
	require.Nil(t, err)
	require.NotNil(t, key)

	k2, err := wutil.LoadKeyFromDER(key.AsDER())
	assert.Nil(t, err)
	assert.NotNil(t, k2)

	pk := key.AsPrivateKey()
	assert.NotNil(t, pk)
	_, ok := pk.(*ecdsa.PrivateKey)
	assert.True(t, ok)

	k3, err := wutil.LoadKeyFromPEM(key.AsPEM())
	assert.Nil(t, err)
	assert.NotNil(t, k3)
}
