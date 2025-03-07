package wutil

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
	"time"

	_ "github.com/go-viper/mapstructure/v2"
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
}

// TlsConfig wraps crypto/tls.Config and adds additional (helper) methods.
type TlsConfig struct {
	// tls.Config is the embedded type of TlsConfig
	tls.Config
	// Profile is a copy of the TlsProfile used to construct this object.  The
	// Cert, Key, and CA fields will always hold PEM data.
	Profile TlsProfile
	// CertFile is the path to a certificate file in PEM format IF the presented
	// TlsProfile contained a file path in the Cert field.
	CertFile string
	// KeyFile is the path to a key file in PEM format IF the presented
	// TlsProfile contained a file path in the Key field.
	KeyFile string
	// CAFile is the path to a CA (bundle) file in PEM format IF the presented
	// TlsProfile Contained a file path in the CA field.
	CAFile string
}

func MakeTlsConfig(profile *TlsProfile) (cfg *TlsConfig, err error) {
	if profile.DisableTls {
		return nil, nil
	}

	// construct a TlsProfile with file contents replacing file paths
	cfg = &TlsConfig{Profile: TlsProfile{
		DisableTls:        profile.DisableTls,
		DisableValidation: profile.DisableValidation,
	}}
	if err = cfg.normalizeToPEMs(profile); err != nil {
		return nil, err
	}

	// optionally load agent certificate and private key
	if len(cfg.Profile.Cert) > 0 {
		cert, err := tls.X509KeyPair([]byte(cfg.Profile.Cert), []byte(cfg.Profile.Key))
		if err != nil {
			return nil, err
		}
		cfg.Certificates = []tls.Certificate{cert}
	}

	// optionally load CA
	if len(cfg.Profile.CA) > 0 {
		caCertPool := x509.NewCertPool()
		if !caCertPool.AppendCertsFromPEM([]byte(cfg.Profile.CA)) {
			return nil, errors.New("x509.AppendCertsFromPem() failed")
		}
		cfg.ClientCAs = caCertPool
		cfg.RootCAs = caCertPool
		cfg.ClientAuth = tls.RequireAndVerifyClientCert // most restrictive
	}

	if profile.DisableValidation {
		cfg.ClientAuth = tls.NoClientCert // disable client cert validation by server
		cfg.InsecureSkipVerify = true     // DANGER - turns off all verification, INSECURE
	}

	return cfg, nil
}

// AllNotBefore scans the entire certificate chain, only the first chain in
// tls.Config.Certificates[], and returns the most restrictive (latest in time)
// value of NotBefore.
func (cfg *TlsConfig) AllNotBefore() (time.Time, error) {
	latest := time.Date(0000, 1, 1, 1, 0, 0, 0, time.UTC) // impossibly far in the past
	if cfg == nil {
		return latest, errors.New("TlsConfig uninitialized (nil)")
	}

	for _, cert := range cfg.Certificates[0].Certificate {
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

// AllNotAfter scans the entire certificate chain, only the first chain in
// tls.Config.Certificates[], and returns the most restrictive (earliest in
// time) value of NotAfter.
func (cfg *TlsConfig) AllNotAfter() (time.Time, error) {
	earliest := time.Date(9999, 12, 31, 23, 59, 59, 9, time.UTC) // impossibly far in the future
	if cfg == nil {
		return earliest, errors.New("TlsConfig uninitalized (nil)")
	}

	for _, cert := range cfg.Certificates[0].Certificate {
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

func (cfg *TlsConfig) normalizeToPEMs(profile *TlsProfile) (err error) {

	// Certificate (chain)
	if cfg.Profile.Cert, err = convertFilePathToPEM(profile.Cert); err != nil {
		return
	}
	if cfg.Profile.Cert != profile.Cert {
		cfg.CertFile = profile.Cert
	}

	// Private Key
	if cfg.Profile.Key, err = convertFilePathToPEM(profile.Key); err != nil {
		return
	}
	if cfg.Profile.Key != profile.Key {
		cfg.KeyFile = profile.Key
	}

	// Certificate Authority (bundle)
	if cfg.Profile.CA, err = convertFilePathToPEM(profile.CA); err != nil {
		return
	}
	if cfg.Profile.CA != profile.CA {
		cfg.CAFile = profile.CA
	}
	return
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
