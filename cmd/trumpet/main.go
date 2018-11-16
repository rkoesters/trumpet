// trumpet is a bot that generates tweets based on the accounts it
// follows.
package main

import (
	"flag"
	"fmt"
	"github.com/rkoesters/trumpet"
	"github.com/rkoesters/trumpet/generator/count"
	"github.com/rkoesters/trumpet/generator/dummy"
	"github.com/rkoesters/trumpet/generator/markov"
	"github.com/rkoesters/trumpet/generator/multi"
	"github.com/rkoesters/trumpet/generator/verbatim"
	"github.com/rkoesters/trumpet/scheduler/noop"
	"github.com/rkoesters/trumpet/scheduler/sametime"
	"github.com/rkoesters/trumpet/scheduler/timer"
	"github.com/rkoesters/trumpet/source/twitter"
	"log"
	"math/rand"
	"os"
	"time"
)

var (
	generator = flag.String("generator", "markov", "name of the generator to use")
	scheduler = flag.String("scheduler", "sametime", "name of the scheduler to use")

	markovLength = flag.Int("markov-length", 3, "length of each prefix for the markov generator")
	timerFreq    = flag.Duration("timer-freq", time.Minute, "frequency for the timer scheduler")

	showVersion = flag.Bool("version", false, "Print version information and exit.")
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

	// Seed the random number generator.
	rand.Seed(time.Now().Unix())

	// Use multi.Generator to multiplex our training data over multiple
	// trumpet.Generators.
	m := multi.New()

	// Use count.Generator to keep track of our input size.
	counter := count.New()
	m.AddTrainer(counter)

	// Use verbatim.Generator to make sure we don't copy a tweet
	// verbatim.
	duplicateChecker := verbatim.New()
	m.AddTrainer(duplicateChecker)

	// Pick our generator.
	gen := pickGenerator()
	m.AddTrainer(gen)
	m.SetGenerator(gen)

	// Pick our scheduler.
	sched := pickScheduler()

	// Prepare the twitter layer.
	err := twitter.Init()
	if err != nil {
		log.Fatal(err)
	}

	// Get list of user IDs to learn from.
	userIDs, err := twitter.GetFriends()
	if err != nil {
		log.Fatal(err)
	}

	// Start fetching our input.
	incomingTweets := make(chan string)
	for _, userID := range userIDs {
		if *scheduler == "sametime" {
			go twitter.GetPastTweets(userID, incomingTweets, noop.New())
		} else {
			go twitter.GetPastTweets(userID, incomingTweets, sched)
		}
	}
	go twitter.ListenForTweets(userIDs, incomingTweets, sched)

	// Start writing tweets according to our scheduler.
	outgoingTweets := composeTweets(m, sched, duplicateChecker)

	// Main loop.
	for {
		select {
		case t := <-incomingTweets:
			log.Printf("IN(((%v)))", t)
			m.Train(t)
			log.Printf("input size: %v", *counter)
		case t := <-outgoingTweets:
			log.Printf("OUT(((%v)))", t)
			err = twitter.Tweet(t)
			if err != nil {
				log.Print("failed posting tweet")
			}
		}
	}
}

func composeTweets(gen trumpet.Generator, sched trumpet.Scheduler, checker *verbatim.Generator) <-chan string {
	c := make(chan string)
	go func() {
		for {
			<-sched.Chan()

			for {
				t := gen.Generate(280)
				if !checker.Exists(t) {
					c <- t
					break
				}
			}
		}
	}()
	return c
}

func pickGenerator() trumpet.Generator {
	switch *generator {
	case "dummy":
		return dummy.New()
	case "markov":
		return markov.New(*markovLength)
	default:
		log.Fatalf("unknown generator: %v", *generator)

		// Won't be reached, but log.Fatal isn't recognized as a
		// valid way to end a function that returns a value.
		panic("log.Fatalf returned")
	}
}

func pickScheduler() trumpet.Scheduler {
	switch *scheduler {
	case "timer":
		return timer.New(*timerFreq)
	case "sametime":
		return sametime.New()
	default:
		log.Fatalf("unknown scheduler: %v", *scheduler)

		// Won't be reached, but log.Fatal isn't recognized as a
		// valid way to end a function that returns a value.
		panic("log.Fatalf returned")
	}
}
