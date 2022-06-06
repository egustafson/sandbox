package main_test

import (
	"strings"
	"testing"

	"github.com/egustafson/sandbox/go/package-dependency-workspace/pkgverlib"
)

func TestDependenciesMain(t *testing.T) {
	versionID := pkgverlib.VersionID()
	if versionID == "-unknown-" {
		t.Errorf("Unknown module version returned: %s", versionID)
	}

	if strings.HasSuffix(versionID, "(devel)") {
		t.Errorf("Development module linked: %s", versionID)
	}
}
