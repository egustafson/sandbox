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

	cachedTlsConfig *tls.Config
	lock            sync.RWMutex
	initAt          time.Time
	certPEM         string
	keyPEM          string
	caPEM           string
}

// TlsConfig returns a *tls.Config constructed from the public fields of the
// TlsProfile.  The *tls.Config is only constructed once.  The cached value is
// returned on subsiquent calls.
func (profile *TlsProfile) TlsConfig() (*tls.Config, error) {
	profile.lock.RLock()

	if profile.cachedTlsConfig != nil || profile.DisableTls {
		defer profile.lock.RUnlock()
		return profile.cachedTlsConfig, nil
	}

	profile.lock.RUnlock()
	return profile.buildTlsConfig()
}

func (profile *TlsProfile) buildTlsConfig() (*tls.Config, error) {
	profile.lock.Lock()
	defer profile.lock.Unlock()

	profile.cachedTlsConfig = &tls.Config{} // initialize with a "nil" configuration
	profile.initAt = time.Now()
	if profile.DisableTls {
		return profile.cachedTlsConfig, nil
	}

	// replace file paths with PEM data if necessary
	if err := profile.loadPEMs(); err != nil { // problem loading files??
		return nil, err
	}

	// optionally load agent certificate and private key
	if len(profile.certPEM) > 0 {
		cert, err := tls.X509KeyPair([]byte(profile.certPEM), []byte(profile.keyPEM))
		if err != nil {
			return nil, err
		}
		// 'cert' is a chain of certificates, config.Certificate is a list of chains
		profile.cachedTlsConfig.Certificates = []tls.Certificate{cert}
	}

	// optionally load a CA
	if len(profile.caPEM) > 0 {
		caCertPool := x509.NewCertPool()
		if !caCertPool.AppendCertsFromPEM([]byte(profile.caPEM)) {
			return nil, errors.New("x509.AppendCertsFromPEM() failed")
		}
		profile.cachedTlsConfig.ClientCAs = caCertPool
		profile.cachedTlsConfig.RootCAs = caCertPool
		profile.cachedTlsConfig.ClientAuth = tls.RequireAndVerifyClientCert // most restrictive
	}

	if profile.DisableValidation {
		profile.cachedTlsConfig.ClientAuth = tls.NoClientCert // disable client cert validation by server
		profile.cachedTlsConfig.InsecureSkipVerify = true     // DANGER - turns off all verification, INSECURE
	}

	return profile.cachedTlsConfig, nil
}

// NotBefore scans the entire certificate chain, only the first chain, in
// tls.Config.Certificates[], and returns the most restrictive (latest in time)
// value of NotBefore.
func (profile *TlsProfile) NotBefore() (time.Time, error) {
	latest := beginningOfTime
	if profile.cachedTlsConfig == nil {
		if profile.DisableTls {
			return latest, errors.New("TlsProfile tls disabled")
		}
		return latest, errors.New("TlsProfile uninitialized")
	}

	for _, cert := range profile.cachedTlsConfig.Certificates[0].Certificate {
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
	if profile.cachedTlsConfig == nil {
		if profile.DisableTls {
			return earliest, errors.New("TlsProfile tls disabled")
		}
		return earliest, errors.New("TlsProfile uninitialized")
	}

	for _, cert := range profile.cachedTlsConfig.Certificates[0].Certificate {
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

func (profile *TlsProfile) loadPEMs() (err error) {
	// Certificate (chain)
	if profile.certPEM, err = convertFilePathToPEM(profile.Cert); err != nil {
		return
	}
	// Private Key
	if profile.keyPEM, err = convertFilePathToPEM(profile.Key); err != nil {
		return
	}
	// Certificate Authority (bundle)
	profile.caPEM, err = convertFilePathToPEM(profile.CA)
	return // profile.{cert,key,ca}PEM now has valid PEM material, or is empty, or returns error
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
