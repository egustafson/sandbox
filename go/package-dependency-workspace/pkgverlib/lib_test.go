package pkgverlib_test

import (
	"fmt"
	"testing"

	"github.com/egustafson/sandbox/go/package-dependency-worksace/pkgverlib"
)

func TestVersionID(t *testing.T) {
	versionID := pkgverlib.VersionID()
	fmt.Printf("Version: %s\n", versionID)
	if versionID != "-unknown-" {
		t.Errorf("Expected ('-unknown-') and received: %s", versionID)
	}
}
