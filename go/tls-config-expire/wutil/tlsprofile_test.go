package wutil_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/egustafson/sandbox/go/tls-config-expire/wutil"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestTLSProfileTestSuite(t *testing.T) {
	suite.Run(t, new(TLSProfileTestSuite))
}

type CertSet struct {
	ca   string
	cert string
	key  string
}

type TLSProfileTestSuite struct {
	suite.Suite
	validFiles         *CertSet
	validPEM           *CertSet
	validSelfSignedPEM *CertSet
	validCaBundlePEM   *CertSet
	validCertChainPEM  *CertSet
	earlyInvalidPEM    *CertSet
	expiredInvalidPEM  *CertSet
}

// SetupAllSuite() at the bottom of this file
//

//
// Necessary tests:
//   1. DisableTLS
//   2. Enabled TLS with all empty PEM blocks
//   3. empty cert/key, populated ca
//   4. empty ca, populated cert/key
//   5. all populated
//   6. all populated + InsecureSkipVerify = true (disable auth)
//
//   7. CA is a bundle
//   8. cert contains a chain of 3

func (s *TLSProfileTestSuite) TestTLSProfile() {
	ca, caKey := wutil.MakeCertAndKey()
	cert, key := wutil.MakeCertAndKey(wutil.WithSigner(ca, caKey))

	tc, err := MakeTempCertFileSet(
		string(ca.AsPEM()),
		string(cert.AsPEM()),
		string(key.AsPEM()),
	)
	require.Nil(s.T(), err)
	defer tc.Delete()

	var profileTests = []struct {
		n string
		p *wutil.TLSProfile
	}{
		{"DisableTLS", &wutil.TLSProfile{DisableTLS: true}},
		{"empty PEMS", &wutil.TLSProfile{ /* enabled tls, empty PEMS */ }},
		{"only CA", &wutil.TLSProfile{
			CA: s.validPEM.ca,
		}},
		{"only cert/key", &wutil.TLSProfile{
			Cert: s.validPEM.cert,
			Key:  s.validPEM.key,
		}},
		{"everything", &wutil.TLSProfile{
			CA:   s.validPEM.ca,
			Cert: s.validPEM.cert,
			Key:  s.validPEM.key,
		}},
		{"everything+InsecureSkipVerify", &wutil.TLSProfile{
			CA:                 s.validPEM.ca,
			Cert:               s.validPEM.cert,
			Key:                s.validPEM.key,
			InsecureSkipVerify: true,
		}},
		{"file ca only", &wutil.TLSProfile{
			CA: s.validFiles.ca,
		}},
		{"file cert/key only", &wutil.TLSProfile{
			Cert: s.validFiles.cert,
			Key:  s.validFiles.key,
		}},
		{"file everything", &wutil.TLSProfile{
			CA:   s.validFiles.ca,
			Cert: s.validFiles.cert,
			Key:  s.validFiles.key,
		}},
	}
	for _, tt := range profileTests {
		s.Run(tt.n, func() {
			_, err := tt.p.ServerConfig()
			require.Nil(s.T(), err)

			s.Nil(tt.p.Validate())
		})
	}
}

func (s *TLSProfileTestSuite) TestTLSProfile_InsecureSkipVerify() {
	TLSProfileInsecureSkipVerify := &wutil.TLSProfile{
		CA:                 s.validPEM.ca,
		Cert:               s.validPEM.cert,
		Key:                s.validPEM.key,
		InsecureSkipVerify: true,
	}
	cfgInsecureSkipVerify, err := TLSProfileInsecureSkipVerify.ServerConfig()
	require.Nil(s.T(), err)
	s.True(cfgInsecureSkipVerify.InsecureSkipVerify)

	TLSProfileYesAuth := &wutil.TLSProfile{
		CA:                 s.validPEM.ca,
		Cert:               s.validPEM.cert,
		Key:                s.validPEM.key,
		InsecureSkipVerify: false, // the default
	}
	cfgVerify, err := TLSProfileYesAuth.ServerConfig()
	require.Nil(s.T(), err)
	s.False(cfgVerify.InsecureSkipVerify)

}

func (s *TLSProfileTestSuite) TestTLSProfile_NonExistantPath() {

	TLSProfile := &wutil.TLSProfile{
		CA:   fmt.Sprintf("bogus-path/%s", s.validFiles.ca),
		Cert: s.validFiles.cert,
		Key:  s.validFiles.key,
	}
	_, err := TLSProfile.ServerConfig()
	s.Error(err)

	profile2 := &wutil.TLSProfile{
		CA:   s.validFiles.ca,
		Cert: "garbage-for-a-cert",
		Key:  s.validFiles.key,
	}
	err = profile2.Validate()
	s.Error(err)

	profile3 := &wutil.TLSProfile{
		CA:   s.validFiles.ca,
		Cert: s.validFiles.cert,
		Key:  "more-garbage",
	}
	err = profile3.Validate()
	s.Error(err)
}

//
// ==  TLSProfileTestSuite -  Setup & TearDown  ===========
//

func (s *TLSProfileTestSuite) SetupSuite() {

	// -- ValidPEM --
	ca, caKey := wutil.MakeCertAndKey()
	cert, key := wutil.MakeCertAndKey(wutil.WithSigner(ca, caKey))
	s.validPEM = &CertSet{
		ca:   string(ca.AsPEM()),
		cert: string(cert.AsPEM()),
		key:  string(key.AsPEM()),
	}

	// -- ValidFiles --
	tc, err := MakeTempCertFileSet(s.validPEM.ca, s.validPEM.cert, s.validPEM.key)
	require.Nil(s.T(), err)
	s.validFiles = &CertSet{
		ca:   tc.caFile.Name(),
		cert: tc.certFile.Name(),
		key:  tc.keyFile.Name(),
	}

	// -- ValidSelfSignedPEM --
	s.validSelfSignedPEM = &CertSet{
		ca:   string(ca.AsPEM()),
		cert: string(ca.AsPEM()),
		key:  string(caKey.AsPEM()),
	}

	// -- ValidCertChainPEM --
	var chainLen uint = 4
	chain := wutil.MakeCertChain(wutil.WithLength(chainLen))
	certChainPEM := ""
	for i := chainLen - 1; i > 0; i-- { // walk leaf -> cert before CA [3..1]
		certChainPEM = fmt.Sprintf("%s\n%s", certChainPEM, string(chain[i].Cert.AsPEM()))
	}
	s.validCertChainPEM = &CertSet{
		ca:   string(chain[0].Cert.AsPEM()),
		cert: certChainPEM,
		key:  string(chain[chainLen-1].Key.AsPEM()),
	}

	// -- ValidCaBundlePEM --
	caBundle := fmt.Sprintf("%s\n%s\n", certChainPEM, ca.AsPEM())
	s.validCaBundlePEM = &CertSet{
		ca:   caBundle,
		cert: string(cert.AsPEM()),
		key:  string(key.AsPEM()),
	}

	// -- EarlyInvalidPEM --
	cert, key = wutil.MakeCertAndKey(
		wutil.WithSigner(ca, caKey),
		wutil.WithNotBefore(time.Now().Add(10*time.Minute)),
	)
	s.earlyInvalidPEM = &CertSet{
		ca:   string(ca.AsPEM()),
		cert: string(cert.AsPEM()),
		key:  string(key.AsPEM()),
	}

	// -- expiredInvalidPEM --
	cert, key = wutil.MakeCertAndKey(
		wutil.WithSigner(ca, caKey),
		wutil.WithNotAfter(time.Now().Add(-time.Minute)),
	)
	s.expiredInvalidPEM = &CertSet{
		ca:   string(ca.AsPEM()),
		cert: string(cert.AsPEM()),
		key:  string(key.AsPEM()),
	}
}

func (s *TLSProfileTestSuite) TearDownSuite() {

	os.Remove(s.validFiles.ca)
	os.Remove(s.validFiles.cert)
	os.ReadDir(s.validFiles.key)
}

// --  Helpers  ----------
//

type TempCerts struct {
	caFile   *os.File
	certFile *os.File
	keyFile  *os.File
}

// MakeTempCertFileSet will create 3 files and populate each of them with the
// contents of the func parameters {ca, cert, key}.  The intent is to pass in
// PEM block(s), although any string, including the empty string will be used.
func MakeTempCertFileSet(ca, cert, key string) (tc *TempCerts, err error) {
	tc = new(TempCerts)
	if tc.caFile, err = os.CreateTemp("", "ca-*.pem"); err != nil {
		return nil, err
	}
	if tc.certFile, err = os.CreateTemp("", "cert-*.pem"); err != nil {
		return nil, err
	}
	if tc.keyFile, err = os.CreateTemp("", "key-*.pem"); err != nil {
		return nil, err
	}

	tc.caFile.WriteString(ca)
	tc.caFile.Close()
	tc.certFile.WriteString(cert)
	tc.certFile.Close()
	tc.keyFile.WriteString(key)
	tc.keyFile.Close()
	return // tc, nil
}

func (tc *TempCerts) Delete() {
	os.Remove(tc.caFile.Name())
	os.Remove(tc.certFile.Name())
	os.Remove(tc.keyFile.Name())
}
