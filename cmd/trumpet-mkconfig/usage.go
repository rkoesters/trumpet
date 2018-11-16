package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	out := flag.CommandLine.Output()

	fmt.Fprintf(out, "Usage: %v [FLAGS...]\n", os.Args[0])
	fmt.Fprintln(out)
	fmt.Fprintln(out, "Flags:")
	flag.PrintDefaults()
}

func init() {
	flag.Usage = usage
}
