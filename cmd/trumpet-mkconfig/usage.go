// trumpet-mkconfig is an interactive trumpet configuration creation
// tool. In addition to creating the configuration file, it also
// verifies that the given credentials are valid and can be used to
// connect to Twitter (this can be disabled with the -skip-verify flag).
//
// For usage information, run:
//
//	$ trumpet-mkconfig -help
package main

import (
	"flag"
	"fmt"
	"os"
)

func init() { flag.Usage = usage }

func usage() {
	fmt.Fprintf(os.Stderr, usageMessage, os.Args[0])
	flag.PrintDefaults()
}

const usageMessage = `Usage: %v [FLAGS...]

Flags:
`
