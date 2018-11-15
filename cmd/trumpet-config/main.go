// trumpet-config is an interactive
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	showVersion = flag.Bool("version", false, "print version information and exit")
	outfile     = flag.String("o", "trumpet.conf", "output file")
)

func main() {
	flag.Parse()

	// We don't take any arguments, only flags.
	if flag.NArg() != 0 {
		flag.Usage()
		os.Exit(1)
	}

	if *showVersion {
		fmt.Println(Version())
		os.Exit(0)
	}

	ckey, csecret, atoken, asecret := getInput()

	writeOutput(ckey, csecret, atoken, asecret)
}

func getInput() (ckey, csecret, atoken, asecret string) {
	sc := bufio.NewScanner(os.Stdin)

	fmt.Print("Consumer Key: ")
	sc.Scan()
	ckey = sc.Text()

	fmt.Print("Consumer Secret: ")
	sc.Scan()
	csecret = sc.Text()

	fmt.Print("Access Token: ")
	sc.Scan()
	atoken = sc.Text()

	fmt.Print("Access Secret: ")
	sc.Scan()
	asecret = sc.Text()

	return
}

func writeOutput(ckey, csecret, atoken, asecret string) {
	f, err := os.Create(*outfile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening %v: %v", *outfile, err)
		os.Exit(1)
	}
	defer f.Close()

	fmt.Fprint(f, "consumer-key=", ckey, "\n")
	fmt.Fprint(f, "consumer-secret=", csecret, "\n")
	fmt.Fprint(f, "access-token=", atoken, "\n")
	fmt.Fprint(f, "access-secret=", asecret, "\n")
}
