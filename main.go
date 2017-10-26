package main

import (
	"flag"
	"log"
	"time"
)

var (
	freq = flag.Duration("freq", time.Hour, "time between tweets")
)

func main() {
	flag.Parse()

	initTwitter()

	gen := &dummyGenerator{}

	incomingTweets := listenForTweets(flag.Args())
	outgoingTweets := composeTweets(gen)

	for {
		select {
		case t := <-incomingTweets:
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
			c <- gen.Generate(140)

			time.Sleep(*freq)
		}
	}()
	return c
}
