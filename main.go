package main

import (
	"flag"
	"github.com/rkoesters/trumpet/gen/markov"
	"log"
	"time"
)

var (
	freq = flag.Duration("freq", time.Hour, "time between tweets")
)

func main() {
	flag.Parse()

	initTwitter()

	userIDs := getUserIDs(flag.Args())

	gen := markov.NewChain(2)
	c := make(chan string)

	for _, userID := range userIDs {
		go getPastTweets(userID, c)
	}
	go listenForTweets(userIDs, c)

	outgoingTweets := composeTweets(gen)

	for {
		select {
		case t := <-c:
			log.Printf("incoming tweet: %v", t)
			gen.Train(t)
		case t := <-outgoingTweets:
			log.Printf("outgoing tweet: %v", t)
			tweet(t)
		}
	}
}

func composeTweets(gen Generator) <-chan string {
	c := make(chan string)
	go func() {
		for {
			time.Sleep(*freq)

			c <- gen.Generate(140)
		}
	}()
	return c
}
