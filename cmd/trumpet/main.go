package main

import (
	"flag"
	"github.com/rkoesters/trumpet"
	"github.com/rkoesters/trumpet/generator/count"
	"github.com/rkoesters/trumpet/generator/markov"
	"github.com/rkoesters/trumpet/generator/multi"
	"github.com/rkoesters/trumpet/generator/verbatim"
	"github.com/rkoesters/trumpet/scheduler/sametime"
	"github.com/rkoesters/trumpet/source/twitter"
	"log"
	"math/rand"
	"os"
	"time"
)

var (
	chainLength = flag.Int("chain-length", 3, "length of each prefix in the markov chain")
)

func main() {
	flag.Parse()

	if flag.NArg() != 0 {
		flag.Usage()
		os.Exit(1)
	}

	rand.Seed(time.Now().Unix())

	twitter.Init()

	userIDs, err := twitter.GetFriends()
	if err != nil {
		log.Fatal(err)
	}

	c := make(chan string)

	gen := multi.New()

	brain := markov.NewChain(*chainLength)
	gen.AddTrainer(brain)
	gen.SetGenerator(brain)

	counter := count.New()
	gen.AddTrainer(counter)

	duplicateChecker := verbatim.New()
	gen.AddTrainer(duplicateChecker)

	sched := sametime.New()

	for _, userID := range userIDs {
		go twitter.GetPastTweets(userID, c)
	}
	go twitter.ListenForTweets(userIDs, c, sched)

	outgoingTweets := composeTweets(gen, sched, duplicateChecker)

	for {
		select {
		case t := <-c:
			log.Printf("IN(((%v)))", t)
			gen.Train(t)
			log.Printf("input size: %v", *counter)
		case t := <-outgoingTweets:
			log.Printf("OUT(((%v)))", t)
			twitter.Tweet(t)
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
