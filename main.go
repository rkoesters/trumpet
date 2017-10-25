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
	initNeuralNetwork()

	incomingTweets := listenForTweets(flag.Args())
	outgoingTweets := composeTweets()

	for {
		select {
		case t := <-incomingTweets:
			log.Printf("incoming tweet: %v", t)
		case t := <-outgoingTweets:
			log.Printf("outgoing tweet: %v", t)
			tweet(t)
		}
	}
}

func composeTweets() <-chan string {
	c := make(chan string)
	go func() {
		for {
			// TODO use neural network to generate tweets
		}
	}()
	return c
}

// TODO remove these in favor of actual functions
func initNeuralNetwork() {}
