// Package sametime provides a trumpet.Scheduler that schedules tweets
// to go out at the same time Train is called.
package sametime

import (
	"time"
)

type Scheduler struct {
	ch chan struct{}
}

func New() *Scheduler {
	return &Scheduler{
		ch: make(chan struct{}),
	}
}

func (s *Scheduler) Train(d time.Time) {
	go func() { s.ch <- struct{}{} }()
}

func (s *Scheduler) Chan() <-chan struct{} {
	return s.ch
}
