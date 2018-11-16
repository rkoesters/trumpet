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
