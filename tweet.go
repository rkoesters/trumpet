package main

import (
	"flag"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"golang.org/x/net/context"
	"log"
)

var (
	consumerKey    = flag.String("consumer-key", "oopq6UwfpkDynEg4Ki6AufyQ3", "")
	consumerSecret = flag.String("consumer-secret", "QBu5KygNsDEsXJb7Q4ZlIObFsAHpkX3dHoNwQxmkDwSYrAkUYu", "")
	accessToken    = flag.String("access-token", "917604710558953474-DBZCatP09K1PF6bjD9I1e7oZfzRTO8s", "")
	accessSecret   = flag.String("access-secret", "MxX9jOMwX0tCegCJYh7SHHLOXM5FHFOth8NL7NozIf7lV", "")
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
