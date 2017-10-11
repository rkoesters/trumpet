package main

import (
	"flag"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"golang.org/x/net/context"
	"log"
)

var (
	consumerKey    = flag.String("consumer-key", "", "")
	consumerSecret = flag.String("consumer-secret", "", "")
	accessToken    = flag.String("access-token", "", "")
	accessSecret   = flag.String("access-secret", "", "")
)

func tweet(s string) {
	config := oauth1.NewConfig(*consumerKey, *consumerSecret)
	token := oauth1.NewToken(*accessToken, *accessSecret)

	client := twitter.NewClient(oauth1.NewClient(context.TODO(), config, token))

	params := &twitter.StatusUpdateParams{
		Status: "hello, world",
	}

	t, _, err := client.Statuses.Update(params.Status, params)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(t.IDStr)
}
