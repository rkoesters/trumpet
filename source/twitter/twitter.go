// Package twitter is the layer that connects trumpet to Twitter so that
// other parts of the program don't need to directly interact with the
// Twitter API.
package twitter

import (
	"flag"
	"github.com/rkoesters/trumpet"
	"github.com/rkoesters/trumpet/source/twitter/logger"
	"github.com/rkoesters/xdg/keyfile"
	"gopkg.in/ChimeraCoder/anaconda.v2"
	"html"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	configFile = flag.String("f", "trumpet.conf", "file to read twitter configuration from")
	live       = flag.Bool("live", false, "send tweets to twitter instead of printing them to stdout")
)

var twitter *anaconda.TwitterApi

// Init prepares the twitter variable for use.
func Init() error {
	file, err := os.Open(*configFile)
	if err != nil {
		return err
	}
	defer file.Close()

	kf, err := keyfile.New(file)
	if err != nil {
		return err
	}

	consumerKey := kf.Value("", "consumer-key")
	consumerSecret := kf.Value("", "consumer-secret")
	accessToken := kf.Value("", "access-token")
	accessSecret := kf.Value("", "access-secret")

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	twitter = anaconda.NewTwitterApi(accessToken, accessSecret)
	twitter.SetLogger(logger.New(logger.LevelInfo))
	ok, err := twitter.VerifyCredentials()
	// ok should be set to false if err != nil, but we are checking
	// both just in case behavior changes.
	if !ok || err != nil {
		log.Print("twitter.VerifyCredentials() failed")
		return err
	}

	return nil
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
func Tweet(s string) error {
	if !*live {
		// silently drop tweet.
		return nil
	}

	_, err := twitter.PostTweet(s, nil)
	return err
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

const PAST_TWEET_REQUESTS = 25

func GetPastTweets(userID string, c chan<- string, sched trumpet.Scheduler) {
	var last string
	for i := 0; i < PAST_TWEET_REQUESTS; i++ {
		v := url.Values{}
		v.Set("user_id", userID)
		v.Set("count", "200")
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
				c <- html.UnescapeString(t.Text)
				last = t.IdStr
				created, err := t.CreatedAtTime()
				if err != nil {
					log.Printf("GetPastTweets: %v", err)
				} else {
					sched.Train(created)
				}
			}
		}
	}
}

// ListenForTweets returns a channel of new tweets posted by the given
// user IDs.
func ListenForTweets(userIDs []string, c chan<- string, sched trumpet.Scheduler) {
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
				c <- html.UnescapeString(msg.Text)
				sched.Train(time.Now())
			}
		}
	}
}
