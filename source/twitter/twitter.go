package twitter

import (
	"flag"
	"github.com/ChimeraCoder/anaconda"
	"log"
	"net/url"
	"strconv"
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

// Init prepares the twitter variable for use.
func Init() {
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

// GetUserIDs takes a slice of twitter user names as input and returns a
// slice of twitter user IDs.
func GetUserIDs(userNames []string) []string {
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

func GetFriends() ([]string, error) {
	var userIDs []string
	var v url.Values

	friendsChan := twitter.GetFriendsIdsAll(v)

	for fidp := range friendsChan {
		for _, id := range fidp.Ids {
			userIDs = append(userIDs, strconv.FormatInt(id, 10))
		}
		if fidp.Error != nil {
			log.Printf("error ranging friends list: %v", fidp.Error)
		}
	}
	return userIDs, nil
}

// Tweet posts a tweet with contents of s.
func Tweet(s string) {
	if !*live {
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

const PAST_TWEET_REQUESTS = 100

func GetPastTweets(userID string, c chan<- string) {
	var last string
	for i := 0; i < PAST_TWEET_REQUESTS; i++ {
		v := url.Values{}
		v.Set("user_id", userID)
		v.Set("count", "50")
		v.Set("exclude_replies", "true")
		if last != "" {
			v.Set("max_id", last)
		}
		timeline, err := twitter.GetUserTimeline(v)
		if err != nil {
			log.Fatal(err)
		}
		for _, t := range timeline {
			if t.IdStr != last && isGoodTweet(t, []string{userID}) {
				c <- t.Text
				last = t.IdStr
			}
		}
	}
}

// ListenForTweets returns a channel of new tweets posted by the given
// user IDs.
func ListenForTweets(userIDs []string, c chan<- string) {
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
