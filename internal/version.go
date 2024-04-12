package internal

import (
	"fmt"
	"runtime"
	"strings"
)

const AppName = "parquet"

// Set via LDFLAGS -X
var (
	Version = "unset"
	Branch  = "unset"
	Commit  = "unset"
)

// parquet/0.1.0 Go/1.20.2
func UserAgent() string {
	return fmt.Sprintf("%s/%s Go/%s",
		AppName, Version, strings.TrimPrefix(runtime.Version(), "go"))
}

// parquet version 0.1.0 (git: main@241913f5) (go: 1.20.2) (os: linux/amd64)
func AppVersion() string {
	return fmt.Sprintf(
		"%s version %s (git: %s@%s) (go: %s) (os: %s/%s)",
		AppName,
		Version,
		Branch,
		Commit,
		strings.TrimLeft(runtime.Version(), "go"),
		runtime.GOOS,
		runtime.GOARCH,
	)
}
