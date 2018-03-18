// Package timer provides a trumpet.Scheduler that schedules tweets
// to go out repeatedly on a timer.
package timer

import (
	"time"
)

type Scheduler struct {
	ch chan struct{}
	d  time.Duration
}

func New(d time.Duration) *Scheduler {
	sched := &Scheduler{
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

func (s *Scheduler) Train(d time.Time) {}

func (s *Scheduler) Chan() <-chan struct{} {
	return s.ch
}
