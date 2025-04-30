package wutil

// The functions in this file are primarily intended to be used in mocking
// tests.  They are included in the package, and not the testing package because
// these functions may be useful beyond the scope of unit testing within this
// package.
//
// WARNING: these functions create certificates, however, there intent is
// focused on creating mock certificates for testing and are likely not hardened
// for production use.

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	crand "crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"math/big"
	mrand "math/rand/v2"
	"net"
	"time"
)

//
// --  Certificate Chain generation  ----------------------
//

type CertKeyPair struct {
	Cert Cert
	Key  Key
}

type CertChain []CertKeyPair

type chainOpts struct {
	CA                    Cert
	CAkey                 Key
	Length                uint
	NotAfterVec           []time.Time
	NotBeforeVec          []time.Time
	IgnoreInconsistencies bool
}

type CertChainOptFn func(*chainOpts)

// MakeCertChain will create a chain of signed certificates and return both the
// certificates and there matching keys.  The first element in the chain is the
// CA, either provided or generated as a self signed Certificate.
func MakeCertChain(opts ...CertChainOptFn) (chain CertChain) {

	const defaultLen = 2

	co := &chainOpts{Length: defaultLen}
	for _, opt := range opts {
		opt(co)
	}

	notBeforeLen := len(co.NotBeforeVec)
	notAfterLen := len(co.NotAfterVec)
	if !co.IgnoreInconsistencies {
		inconsistent := ((notBeforeLen > 0) && (notAfterLen > 0) && (notBeforeLen != notAfterLen)) ||
			((co.Length != defaultLen) && (notBeforeLen > 0) && (notBeforeLen != int(co.Length))) ||
			((co.Length != defaultLen) && (notAfterLen > 0) && (notAfterLen != int(co.Length)))
		if inconsistent {
			panic("chain length and not before/after vector lengths are inconsistent")
		}
		caStated := !co.CA.IsNil() || !co.CAkey.IsNil()
		if caStated && (co.CA.IsNil() || co.CAkey.IsNil()) {
			panic("both CA cert and CA key must be given, or neither")
		}
	}

	if notBeforeLen > 0 || notAfterLen > 0 { // extend co.Length if NotBefore/After is specified
		if notBeforeLen > notAfterLen {
			co.Length = uint(notBeforeLen)
		} else {
			co.Length = uint(notAfterLen)
		}
	}

	// now generate the Cert/Key pairs into a chain

	idx := 0
	chain = make([]CertKeyPair, 0, co.Length)
	if co.CA != nil {
		chain = append(chain, CertKeyPair{Cert: co.CA, Key: co.CAkey})
		idx += 1
	}

	for idx < int(co.Length) {
		certOps := make([]CertOptFn, 0)
		if idx > 0 { // sign with previous cert
			certOps = append(certOps, WithSigner(chain[idx-1].Cert, chain[idx-1].Key))
		}
		if len(co.NotBeforeVec) > idx {
			certOps = append(certOps, WithNotBefore(co.NotBeforeVec[idx]))
		}
		if len(co.NotAfterVec) > idx {
			certOps = append(certOps, WithNotAfter(co.NotAfterVec[idx]))
		}

		cert, key := MakeCertAndKey(certOps...)
		chain = append(chain, CertKeyPair{Cert: cert, Key: key})
		idx += 1
	}
	return
}

func WithCA(caCert Cert, caKey Key) CertChainOptFn {
	return func(co *chainOpts) {
		co.CA = caCert
		co.CAkey = caKey
	}
}

func WithLength(length uint) CertChainOptFn {
	return func(co *chainOpts) {
		co.Length = length
	}
}

func WithNotAfterVec(notAfters []time.Time) CertChainOptFn {
	return func(co *chainOpts) {
		co.NotAfterVec = notAfters
	}
}

func WithNotBeforeVec(notBefores []time.Time) CertChainOptFn {
	return func(co *chainOpts) {
		co.NotBeforeVec = notBefores
	}
}

func WithIgnoreInconsistencies(f bool) CertChainOptFn {
	return func(co *chainOpts) {
		co.IgnoreInconsistencies = f
	}
}

func (c CertChain) Leaf() CertKeyPair {
	if len(c) > 0 {
		return c[len(c)-1]
	}
	return CertKeyPair{}
}

func (c CertChain) CA() CertKeyPair {
	if len(c) > 0 {
		return c[0]
	}
	return CertKeyPair{}
}

func (c CertChain) Intermediates() CertChain {
	if len(c) > 2 {
		left := 1
		right := len(c) - 1
		return c[left:right]
	}
	return CertChain{} // empty list
}

//
// --  Certificate generation  ----------------------------
//

type certProfile struct {
	template  *x509.Certificate
	signer    *x509.Certificate
	signerKey any
}

type CertOptFn func(*certProfile)

// MakeCertAndKey returns a X.509 Certificate and matching Key.  There are
// options to provide many of the configurable aspects of the certificate,
// including signing the certificate with a parent certificate.  By default the
// returned Certificate is self-signed and good for 1 hour before and 1 hour
// after the current time.
func MakeCertAndKey(opts ...CertOptFn) (cert Cert, key Key) {

	sn := big.NewInt(mrand.Int64())
	template := x509.Certificate{ // define the default Certificate for maximum usability
		SerialNumber:          sn,
		Subject:               pkix.Name{CommonName: fmt.Sprintf("cert-sn-%s", sn.String())},
		NotBefore:             time.Now().Add(-time.Hour), // 1 hour before now
		NotAfter:              time.Now().Add(time.Hour),  // 1 hour after now
		KeyUsage:              x509.KeyUsage(511),         // enable all usage
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
		BasicConstraintsValid: true,
		IsCA:                  true,
		DNSNames:              []string{"localhost"},
	}

	// FUTURE: configurable crypto algorithms, Elliptic Curve is hard coded here.
	//
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), crand.Reader)
	if err != nil {
		panic(err)
	}

	cp := certProfile{
		template:  &template,
		signer:    &template,  // defaults to self signed
		signerKey: privateKey, // defaults to self signed
	}

	for _, opt := range opts { // apply the options
		opt(&cp)
	}

	certBytes, err := x509.CreateCertificate(
		crand.Reader,          // randomness
		cp.template,           // template to form certificate from
		cp.signer,             // signer certificate
		&privateKey.PublicKey, // this (new) certificate's public key
		cp.signerKey,          // signer private key
	)
	if err != nil {
		panic(err)
	}
	cert = Cert(certBytes) // from x509.go in wutil
	key = Key{privateKey}  // from x509.go in wutil

	// intended for use when unit testing and as an example of how to emit PEMs
	const debug = false
	if debug {
		fmt.Print(string(cert.AsPEM()))
		fmt.Print(string(key.AsPEM()))
	} // end debug

	return // cert, key
}

func WithSigner(signer Cert, signerKey Key) CertOptFn {
	return func(cp *certProfile) {
		cp.signer = signer.AsX509Certificate()
		cp.signerKey = signerKey.AsPrivateKey()
	}
}

func WithNotBefore(t time.Time) CertOptFn {
	return func(cp *certProfile) {
		cp.template.NotBefore = t
	}
}

func WithNotAfter(t time.Time) CertOptFn {
	return func(cp *certProfile) {
		cp.template.NotAfter = t
	}
}

func WithSubject(rdn pkix.Name) CertOptFn {
	return func(cp *certProfile) {
		cp.template.Subject = rdn
	}
}

func WithKeyUsage(usage x509.KeyUsage) CertOptFn {
	return func(cp *certProfile) {
		cp.template.KeyUsage = usage
	}
}

func WithExtKeyUsage(extUsage []x509.ExtKeyUsage) CertOptFn {
	return func(cp *certProfile) {
		cp.template.ExtKeyUsage = extUsage
	}
}

func WithSerialNumber(sn *big.Int) CertOptFn {
	return func(cp *certProfile) {
		cp.template.SerialNumber = sn
	}
}

func WithBasicConstraintsValid(f bool) CertOptFn {
	return func(cp *certProfile) {
		cp.template.BasicConstraintsValid = f
	}
}

func WithIsCA(f bool) CertOptFn {
	return func(cp *certProfile) {
		cp.template.BasicConstraintsValid = true
		cp.template.IsCA = f
	}
}

func WithMaxPathLen(pathLen int) CertOptFn {
	return func(cp *certProfile) {
		cp.template.BasicConstraintsValid = true
		cp.template.MaxPathLenZero = true
		cp.template.MaxPathLen = pathLen
	}
}

func WithDNSNames(names []string) CertOptFn {
	return func(cp *certProfile) {
		cp.template.DNSNames = names
	}
}

func WithIPAddresses(addrs []net.IP) CertOptFn {
	return func(cp *certProfile) {
		cp.template.IPAddresses = addrs
	}
}
