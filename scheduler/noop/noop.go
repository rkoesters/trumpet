// Package noop provides a trumpet.Scheduler that does nothing.
package noop

import (
	"github.com/rkoesters/trumpet"
	"time"
)

type scheduler struct {
	ch chan struct{}
}

// New returns a trumpet.Scheduler that doesn't do anything. It returns
// a valid channel, but nothing will ever be sent over that channel.
func New() trumpet.Scheduler {
	return &scheduler{
		ch: make(chan struct{}),
	}
}

func (g *scheduler) Train(d time.Time) {}

func (g *scheduler) Chan() <-chan struct{} { return g.ch }
