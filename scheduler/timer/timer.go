// Package timer provides a trumpet.Scheduler that schedules tweets
// to go out repeatedly on a timer.
package timer

import (
	"github.com/rkoesters/trumpet"
	"time"
)

type scheduler struct {
	ch chan struct{}
	d  time.Duration
}

// New returns a trumpet.Scheduler that schedules a tweet for the given
// time.Duration after each <-Chan().
func New(d time.Duration) trumpet.Scheduler {
	sched := &scheduler{
		ch: make(chan struct{}),
		d:  d,
	}

	go func() {
		for {
			time.Sleep(sched.d)
			sched.ch <- struct{}{}
		}
	}()

	return sched
}

func (s *scheduler) Train(d time.Time) {}

func (s *scheduler) Chan() <-chan struct{} {
	return s.ch
}
