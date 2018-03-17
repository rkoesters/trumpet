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
