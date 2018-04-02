// Package noop provides a trumpet.Scheduler that does nothing.
package noop

import (
	"time"
)

type Generator struct {
	ch chan struct{}
}

func New() *Generator {
	return &Generator{
		ch: make(chan struct{}),
	}
}

func (g *Generator) Train(d time.Time) {}

func (g *Generator) Chan() <-chan struct{} { return g.ch }
