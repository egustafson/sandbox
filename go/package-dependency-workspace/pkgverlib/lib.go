package pkgverlib

import (
	"fmt"
	"log"
	"runtime/debug"
)

const (
	SelfPackageName = "github.com/egustafson/sandbox/go/pkgverlib"
)

// VersionID returns a string that identifies the package version.
func VersionID() string {

	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		log.Printf("Failed to read build info")
		return "-error-"
	}

	for _, module := range buildInfo.Deps {
		if module.Path == SelfPackageName {
			return fmt.Sprintf("%s@%s", module.Path, module.Version)
		}
	}

	return "-unknown-"
}

func BuildSettings() map[string]string {
	settings := make(map[string]string)

	return settings
}
