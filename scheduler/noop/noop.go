// Package noop provides a trumpet.Scheduler that does nothing.
package noop

import (
	"time"
)

type generator struct {
	ch chan struct{}
}

// New returns a trumpet.Scheduler that doesn't do anything. It returns
// a valid channel, but nothing will ever be sent over that channel.
func New() trumpet.Generator {
	return &generator{
		ch: make(chan struct{}),
	}
}

func (g *generator) Train(d time.Time) {}

func (g *generator) Chan() <-chan struct{} { return g.ch }
