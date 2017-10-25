package main

import (
	"flag"
	"github.com/ChimeraCoder/anaconda"
	"log"
	"net/url"
	"strings"
)

var (
	consumerKey    = flag.String("consumer-key", "oopq6UwfpkDynEg4Ki6AufyQ3", "twitter consumer key")
	consumerSecret = flag.String("consumer-secret", "QBu5KygNsDEsXJb7Q4ZlIObFsAHpkX3dHoNwQxmkDwSYrAkUYu", "twitter consumer secret")
	accessToken    = flag.String("access-token", "917604710558953474-DBZCatP09K1PF6bjD9I1e7oZfzRTO8s", "twitter access token")
	accessSecret   = flag.String("access-secret", "MxX9jOMwX0tCegCJYh7SHHLOXM5FHFOth8NL7NozIf7lV", "twitter access secret")

	live = flag.Bool("live", false, "send tweets to twitter instead of printing them to stdout")
)

var twitter *anaconda.TwitterApi

func initTwitter() {
	anaconda.SetConsumerKey(*consumerKey)
	anaconda.SetConsumerSecret(*consumerSecret)
	twitter = anaconda.NewTwitterApi(*accessToken, *accessSecret)
	twitter.SetLogger(anaconda.BasicLogger)
}

// tweet posts a tweet with contents s
func tweet(s string) {
	if !*live {
		log.Printf("tweet: %s", s)
		return
	}

	_, err := twitter.PostTweet(s, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func listenForTweets(following []string) <-chan string {
	c := make(chan string)
	go func() {
		var userIds []string

		// get IDs for users in following
		for _, userName := range following {
			u, err := twitter.GetUsersShow(userName, nil)
			if err != nil {
				log.Fatal(err)
			}
			userIds = append(userIds, u.IdStr)
		}

		// start listening for tweets from twitter
		v := url.Values{}
		v.Set("follow", strings.Join(userIds, ","))
		stream := twitter.PublicStreamFilter(v)
		defer stream.Stop()

		// iterate over incoming messages from twitter
		for i := range stream.C {
			switch msg := i.(type) {
			default:
				log.Printf("unknown message: %v", msg)
			case anaconda.Tweet:
				// we only want tweets created by the
				// users in userIds
				if isStringInSlice(msg.User.IdStr, userIds) {
					c <- msg.Text
				}
			}
		}
	}()
	return c
}
