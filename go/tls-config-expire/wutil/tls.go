package wutil

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
	"sync"
	"time"

	_ "github.com/go-viper/mapstructure/v2"
)

var (
	beginningOfTime = time.Date(0000, 1, 1, 1, 0, 0, 0, time.UTC)      // impossibly far in the past
	endOfTime       = time.Date(9999, 12, 31, 23, 59, 59, 9, time.UTC) // impossibly far in the future
)

// TlsProfile holds the raw input used to build a TlsConfig, (which is an
// extension of a crypto/tls.Config).
type TlsProfile struct {
	// Cert holds either the path to, or the complete PEM data of the
	// certificate, including all necessary intermediaries.  The leaf
	// certificate must be the first PEM block in either the file or the data.
	Cert string `mapstructure:"cert,omitempty"`
	// Key holds either the path to, or PEM data for the private key paired with
	// `Cert`.
	Key string `mapstructure:"key,omitempty"`
	// CA holds either the path to, or the complete PEM data of certificates to
	// be considered authoritative signers of certificates.  If the CA field is
	// an empty string then the ClientCAs and RootCAs fields in tls.Config will
	// be null.
	CA string `mapstructure:"ca,omitempty"`
	// DisableTls is a flag that indicates TLS would not be used, it is mostly
	// informational.  If DisableTls is set then the Cert, Key, and CA fields
	// may be empty and they will not be processed by functions in this module.
	DisableTls bool `mapstructure:"disable-tls,omitempty"`
	// DisableValidation is a flag that indicates that connections with remote
	// agents should not be validated against the CA and the CA field may be
	// blank.
	DisableValidation bool `mapstructure:"disable-validation,omitempty"`
	// details holds the internal details of the TlsProfile object.
	details tlsDetails
}

// tlsDetails holds the hidden, implementation details of the TlsProfile.
type tlsDetails struct {
	config *tls.Config
	lock   sync.RWMutex
	initAt time.Time
	// certFilePath will hold the filepath IIF TlsProfile.Cert was initialized with a filepath.
	certFilePath string
	// keyFilePath will hold the filepath IIF TlsProfile.Key was initialized with a filepath.
	keyFilePath string
	// caFilePath will hold the filepath IIF TlsProfile.CA was initialized with a filepath.
	caFilePath string
}

// TlsConfig returns a *tls.Config constructed from the public fields of the
// TlsProfile.  The *tls.Config is only constructed once.  The cached value is
// returned on subsiquent calls.
func (profile *TlsProfile) TlsConfig() (*tls.Config, error) {
	profile.details.lock.RLock()

	if profile.details.config != nil || profile.DisableTls {
		defer profile.details.lock.RUnlock()
		return profile.details.config, nil
	}

	profile.details.lock.RUnlock()
	return profile.buildTlsConfig()
}

func (profile *TlsProfile) buildTlsConfig() (*tls.Config, error) {
	profile.details.lock.Lock()
	defer profile.details.lock.Unlock()

	profile.details.config = &tls.Config{} // initialize with a "nil" configuration
	profile.details.initAt = time.Now()
	if profile.DisableTls {
		return profile.details.config, nil
	}

	// replace file paths with PEM data if necessary
	if err := profile.normalizeToPEMs(); err != nil { // problem loading files??
		return nil, err
	}

	// optionally load agent certificate and private key
	if len(profile.Cert) > 0 {
		cert, err := tls.X509KeyPair([]byte(profile.Cert), []byte(profile.Key))
		if err != nil {
			return nil, err
		}
		// 'cert' is a chain of certificates, config.Certificate is a list of chains
		profile.details.config.Certificates = []tls.Certificate{cert}
	}

	// optionally load a CA
	if len(profile.CA) > 0 {
		caCertPool := x509.NewCertPool()
		if !caCertPool.AppendCertsFromPEM([]byte(profile.CA)) {
			return nil, errors.New("x509.AppendCertsFromPEM() failed")
		}
		profile.details.config.ClientCAs = caCertPool
		profile.details.config.RootCAs = caCertPool
		profile.details.config.ClientAuth = tls.RequireAndVerifyClientCert // most restrictive
	}

	if profile.DisableValidation {
		profile.details.config.ClientAuth = tls.NoClientCert // disable client cert validation by server
		profile.details.config.InsecureSkipVerify = true     // DANGER - turns off all verification, INSECURE
	}

	return profile.details.config, nil
}

// NotBefore scans the entire certificate chain, only the first chain, in
// tls.Config.Certificates[], and returns the most restrictive (latest in time)
// value of NotBefore.
func (profile *TlsProfile) NotBefore() (time.Time, error) {
	latest := beginningOfTime
	if profile.details.config == nil {
		if profile.DisableTls {
			return latest, errors.New("TlsProfile tls disabled")
		}
		return latest, errors.New("TlsProfile uninitialized")
	}

	for _, cert := range profile.details.config.Certificates[0].Certificate {
		x509cert, err := x509.ParseCertificate(cert)
		if err != nil {
			return latest, err
		}
		if latest.Before(x509cert.NotBefore) {
			latest = x509cert.NotBefore
		}
	}
	return latest, nil
}

// NotAfter scans the entire certificate chain, only the first chain, in
// tls.Config.Certificates[], and returns the most restrictive (earliest in
// time) value of NotAfter.
func (profile *TlsProfile) NotAfter() (time.Time, error) {
	earliest := endOfTime
	if profile.details.config == nil {
		if profile.DisableTls {
			return earliest, errors.New("TlsProfile tls disabled")
		}
		return earliest, errors.New("TlsProfile uninitialized")
	}

	for _, cert := range profile.details.config.Certificates[0].Certificate {
		x509cert, err := x509.ParseCertificate(cert)
		if err != nil {
			return earliest, err
		}
		if earliest.After(x509cert.NotAfter) {
			earliest = x509cert.NotAfter
		}
	}
	return earliest, nil
}

// normalizeToPEMs will ensure that the public TlsProfile fields (Cert, Key, CA)
// hold PEM data.  If those fields are initialized with file paths then the file
// paths will be saved in details.xxxFilePath fields within the details struct.
func (profile *TlsProfile) normalizeToPEMs() (err error) {

	profile.details.certFilePath = profile.Cert
	profile.details.keyFilePath = profile.Key
	profile.details.caFilePath = profile.CA

	// Certificate (chain)
	if profile.Cert, err = convertFilePathToPEM(profile.Cert); err != nil {
		return
	}
	if profile.details.certFilePath == profile.Cert { // the data was a PEM, not a file path
		profile.details.certFilePath = ""
	}

	// Private Key
	if profile.Key, err = convertFilePathToPEM(profile.Key); err != nil {
		return
	}
	if profile.details.keyFilePath == profile.Key { // the data was a PEM, not a file path
		profile.details.keyFilePath = ""
	}

	// Certificate Authority (bundle)
	if profile.CA, err = convertFilePathToPEM(profile.CA); err != nil {
		return
	}
	if profile.details.caFilePath == profile.CA { // the data was a PEM, not a file path
		profile.details.caFilePath = ""
	}
	return // profile.{Cert,Key,CA} now has valid PEM material.
}

// convertFilePathToPEM attempts to open `path` and read the contents from the
// file.  If successful, the contents of the file are returned as the result.
// If unsuccessful, the original string, `path` is returned.
func convertFilePathToPEM(path string) (string, error) {
	if len(path) < 1 {
		return path, nil
	}

	// check if a PEM block is in `path`
	_, rest := pem.Decode([]byte(path))
	if len(rest) < len(path) {
		return path, nil
	}

	data, err := os.ReadFile(path)
	return string(data), err
}
