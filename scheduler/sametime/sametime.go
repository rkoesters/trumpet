// Package sametime provides a trumpet.Scheduler that schedules tweets
// to go out at the same time Train is called.
package sametime

import (
	"github.com/rkoesters/trumpet"
	"time"
)

type scheduler struct {
	ch chan struct{}
}

// New returns a trumpet.Scheduler that schedules tweets to go out at
// the same time Train is called.
func New() trumpet.Scheduler {
	return &scheduler{
		ch: make(chan struct{}),
	}
}

func (s *scheduler) Train(d time.Time) {
	go func() { s.ch <- struct{}{} }()
}

func (s *scheduler) Chan() <-chan struct{} {
	return s.ch
}
