package main_test

import (
	"testing"

	"github.com/egustafson/sandbox/go/tls-config-expire/wutil"
)

func TestAsTlsConfig(t *testing.T) {

	tlsProfile := &wutil.TlsProfile{
		DisableTls: true,
	}

	tlsConfig, err := wutil.MakeTlsConfig(tlsProfile)
	if err != nil {
		t.FailNow()
	}

	_ = tlsConfig.Clone()
}
