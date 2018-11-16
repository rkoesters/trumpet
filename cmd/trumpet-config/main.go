// trumpet-config is an interactive trumpet configuration creation tool.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/rkoesters/trumpet/source/twitter/logger"
	"gopkg.in/ChimeraCoder/anaconda.v2"
	"os"
	"strings"
)

var (
	showVersion = flag.Bool("version", false, "print version information and exit")
	outfile     = flag.String("o", "trumpet.conf", "output file")
	skipVerify  = flag.Bool("skip-verify", false, "skip credential verification")
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

	if !*skipVerify {
		verifyCredentials(ckey, csecret, atoken, asecret)
	}

	writeOutput(ckey, csecret, atoken, asecret)

	fmt.Println("Wrote configuration to", *outfile)
}

func getInput() (ckey, csecret, atoken, asecret string) {
	sc := bufio.NewScanner(os.Stdin)

	fmt.Print("Consumer Key: ")
	sc.Scan()
	ckey = strings.TrimSpace(sc.Text())

	fmt.Print("Consumer Secret: ")
	sc.Scan()
	csecret = strings.TrimSpace(sc.Text())

	fmt.Print("Access Token: ")
	sc.Scan()
	atoken = strings.TrimSpace(sc.Text())

	fmt.Print("Access Secret: ")
	sc.Scan()
	asecret = strings.TrimSpace(sc.Text())

	return
}

func verifyCredentials(ckey, csecret, atoken, asecret string) {
	anaconda.SetConsumerKey(ckey)
	anaconda.SetConsumerSecret(csecret)

	tapi := anaconda.NewTwitterApi(atoken, asecret)
	tapi.SetLogger(logger.New(logger.LevelInfo))

	ok, _ := tapi.VerifyCredentials()

	if !ok {
		fmt.Fprintf(os.Stderr, "Failed to verify credentials\n")
		os.Exit(1)
	}
}

func writeOutput(ckey, csecret, atoken, asecret string) {
	f, err := os.Create(*outfile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening %v: %v\n", *outfile, err)
		os.Exit(1)
	}
	defer f.Close()

	fmt.Fprint(f, "consumer-key=", ckey, "\n")
	fmt.Fprint(f, "consumer-secret=", csecret, "\n")
	fmt.Fprint(f, "access-token=", atoken, "\n")
	fmt.Fprint(f, "access-secret=", asecret, "\n")
}
