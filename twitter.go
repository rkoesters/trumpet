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

// initTwitter prepares the twitter variable for use.
func initTwitter() {
	anaconda.SetConsumerKey(*consumerKey)
	anaconda.SetConsumerSecret(*consumerSecret)
	twitter = anaconda.NewTwitterApi(*accessToken, *accessSecret)
	twitter.SetLogger(anaconda.BasicLogger)
	ok, err := twitter.VerifyCredentials()
	// ok should be set to false if err != nil, but we are checking
	// both just in case behavior changes.
	if !ok || err != nil {
		log.Print("twitter.VerifyCredentials() failed")
		log.Fatal(err)
	}
}

// getUserIDs takes a slice of twitter user names as input and returns a
// slice of twitter user IDs.
func getUserIDs(userNames []string) []string {
	var userIDs []string

	for _, userName := range userNames {
		u, err := twitter.GetUsersShow(userName, nil)
		if err != nil {
			log.Fatal(err)
		}
		userIDs = append(userIDs, u.IdStr)
	}
	return userIDs
}

// tweet posts a tweet with contents of s.
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

// isGoodTweet returns true if the tweet was written by one of the users
// specified in userIDs.
func isGoodTweet(t anaconda.Tweet, userIDs []string) bool {
	if t.RetweetedStatus != nil {
		// user didn't write tweet, they retweeted someone else.
		return false
	}
	if !isStringInSlice(t.User.IdStr, userIDs) {
		// tweet wasn't made by someone in userIDs.
		return false
	}
	return true
}

func getPastTweets(userID string, c chan<- string) {
	v := url.Values{}
	v.Set("user_id", userID)
	timeline, err := twitter.GetUserTimeline(v)
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range timeline {
		if isGoodTweet(t, []string{userID}) {
			c <- t.Text
		}
	}
}

// listenForTweets returns a channel of new tweets posted by the given
// user IDs.
func listenForTweets(userIDs []string, c chan<- string) {
	// start listening for tweets from twitter
	v := url.Values{}
	v.Set("follow", strings.Join(userIDs, ","))
	stream := twitter.PublicStreamFilter(v)
	defer stream.Stop()

	// iterate over incoming messages from twitter
	for i := range stream.C {
		switch msg := i.(type) {
		default:
			log.Printf("unknown message: %v", msg)
		case anaconda.Tweet:
			if isGoodTweet(msg, userIDs) {
				c <- msg.Text
			}
		}
	}
}
