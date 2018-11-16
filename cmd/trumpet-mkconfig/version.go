package main

import (
	"fmt"
	"runtime"
)

var version = "undefined"

// Version returns the version string for this binary.
func Version() string {
	return fmt.Sprintf(
		"trumpet-mkconfig version %v (%v %v/%v)",
		version,
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH,
	)
}
