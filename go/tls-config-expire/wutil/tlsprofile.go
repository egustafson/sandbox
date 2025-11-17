package wutil

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
	"time"
)

var (
	beginningOfTime = time.Date(0000, 1, 1, 1, 0, 0, 0, time.UTC)      // impossibly far in the past
	endOfTime       = time.Date(9999, 12, 31, 23, 59, 59, 9, time.UTC) // impossibly far in the future
)

// TLSProfile represents a TLS configuration profile.  It is usable for both
// client and server TLS configurations.
type TLSProfile struct {
	// Cert holds either a (chain) of PEM block(s) or a PEM file path.
	Cert string `yaml:"cert" json:"cert" mapstructure:"cert"`

	// Key holds either a PEM block or a path to a PEM file.
	Key string `yaml:"key" json:"key" mapstructure:"key"`

	// CA holds either a (bundle) of PEM block(s) or PEM file path.
	CA string `yaml:"ca" json:"ca" mapstructure:"ca"`

	// DisableTLS indicates that TLS should be disabled. (all other fields are
	// then ignored.)
	DisableTLS bool `yaml:"disable-tls" json:"disable-tls" mapstructure:"disable-tls"`

	// InsecureSkipVerify indicates whether to skip verification of the server's
	// certificate chain and host name.  This flag is only used by clients.
	// Hardened clients do not skip verification of a server's cert.
	InsecureSkipVerify bool `yaml:"insecure-skip-verify,omitempty" json:"insecure-skip-verify,omitempty" mapstructure:"insecure-skip-verify,omitempty"`

	// ClientVerify (mTLS) indicates whether the server should request and
	// verify client certificates.  This flag is only used by servers.
	ClientVerify bool `yaml:"client-verify,omitempty" json:"client-verify,omitempty" mapstructure:"client-verify,omitempty"`
}

// IsTLSDisabled returns true if TLS is disabled in the profile.
func (p *TLSProfile) IsTLSDisabled() bool {
	return p != nil && p.DisableTLS
}

// Validate checks the TLSProfile for validity, specifically checking the
// NotBefore and NotAfter times in the certificate (chain), and the loadability
// of the {CA,Cert,Key} fields.
func (p *TLSProfile) Validate() error {
	if p.DisableTLS {
		return nil // disabled profiles are always valid.
	}

	// build the tls.Config -> ensures CAs load properly
	tlsConfig, err := p.buildTLSConfig()
	if err != nil {
		return err
	}

	// verify the GetCertificate function is set
	if tlsConfig.GetCertificate == nil {
		return errors.New("GetCertificate function is nul in TLS config")
	}

	// confirm the cert (chain) is neither expired nor not yet valid
	notBefore, notAfter := p.effectiveTime()
	now := time.Now()
	if now.Before(notBefore) {
		return errors.New("certificate is not yet valid")
	}
	if now.After(notAfter) {
		return errors.New("certificate has expired")
	}
	return nil
}

func (p *TLSProfile) effectiveTime() (time.Time, time.Time) {
	cert, err := p.loadCert()
	if err != nil {
		return endOfTime, beginningOfTime // invalid cert; return impossible times
	}
	if cert == nil || len(cert.Certificate) == 0 {
		return beginningOfTime, endOfTime // no cert; valid for all time
	}
	notAfter := endOfTime
	notBefore := beginningOfTime
	for _, cert := range cert.Certificate {
		x509Cert, err := x509.ParseCertificate(cert)
		if err != nil {
			return endOfTime, beginningOfTime // indicate an error; return impossible times
		}
		if x509Cert.NotBefore.After(notBefore) {
			notBefore = x509Cert.NotBefore
		}
		if x509Cert.NotAfter.Before(notAfter) {
			notAfter = x509Cert.NotAfter
		}
	}
	return notBefore, notAfter
}

// ServerConfig builds a tls.Config for server use.
func (p *TLSProfile) ServerConfig() (*tls.Config, error) {

	tlsConfig, err := p.buildTLSConfig()
	if err != nil {
		return nil, err
	}
	tlsConfig.InsecureSkipVerify = p.InsecureSkipVerify
	return tlsConfig, nil
}

// ClientConfig builds a tls.Config for client use.
func (p *TLSProfile) ClientConfig() (*tls.Config, error) {

	tlsConfig, err := p.buildTLSConfig()
	if err != nil {
		return nil, err
	}
	if p.ClientVerify { // mTLS
		tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
	}
	return tlsConfig, nil
}

// buildTLSConfig builds the base TLSConfig for both client and server use.
func (p *TLSProfile) buildTLSConfig() (*tls.Config, error) {

	tlsConfig := &tls.Config{} // initialize with an empty config
	if p.DisableTLS {
		return tlsConfig, nil // an empty config is valid
	}

	// dynamically load certs/keys/CAs on each handshake
	tlsConfig.GetCertificate = func(_ *tls.ClientHelloInfo) (*tls.Certificate, error) {
		return p.loadCert()
	}

	caPool, err := p.loadCAPool()
	if err != nil {
		return nil, err
	}
	if caPool != nil {
		tlsConfig.ClientCAs = caPool
		tlsConfig.RootCAs = caPool
	}
	return tlsConfig, nil
}

// loadCert uses loadFileOrPem() to load the certificate and key.
func (p *TLSProfile) loadCert() (*tls.Certificate, error) {
	certPEM, err := loadFileOrPEM(p.Cert)
	if err != nil {
		return nil, err
	}
	keyPEM, err := loadFileOrPEM(p.Key)
	if err != nil {
		return nil, err
	}
	if len(certPEM) < 1 {
		return nil, nil // an empty cert is valid
	}
	cert, err := tls.X509KeyPair([]byte(certPEM), []byte(keyPEM))
	if err != nil {
		return nil, err
	}
	return &cert, nil
}

// loadCAPool loads the CA certificate pool from the profile using
// loadFileOrPEM().
func (p *TLSProfile) loadCAPool() (*x509.CertPool, error) {
	if p.CA == "" {
		return nil, nil // an empty CA is valid
	}
	caPEM, err := loadFileOrPEM(p.CA)
	if err != nil {
		return nil, err
	}
	caPool := x509.NewCertPool()
	if ok := caPool.AppendCertsFromPEM([]byte(caPEM)); !ok {
		return nil, errors.New("failed to parse CA certificate(s)")
	}
	return caPool, nil
}

// loadFileOrPEM loads either a PEM string or a file containing PEM data.
func loadFileOrPEM(input string) (string, error) {
	if input == "" {
		return "", nil // an empty PEM is valid
	}

	_, rest := pem.Decode([]byte(input))
	if len(rest) < len(input) { // ==> there is valid PEM data in input
		// multiple blocks may be present; return the whole input string
		return input, nil
	}

	// else, treat input as a file path
	data, err := os.ReadFile(input)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
