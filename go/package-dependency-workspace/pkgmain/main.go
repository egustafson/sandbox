package main

import (
	"fmt"

	"github.com/egustafson/sandbox/go/package-dependency-workspace/pkgverlib"
)

func main() {
	versionID := pkgverlib.VersionID()
	fmt.Printf("Package Version: %s\n", versionID)
}
