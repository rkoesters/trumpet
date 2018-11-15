// Package dummy provides a dummy trumpet.Generator.
package dummy

import (
	"github.com/rkoesters/trumpet"
)

type generator struct{}

// New returns a trumpet.Generator that ignores training data and only
// generates the string "hello, world" truncated to the requested
// length.
func New() trumpet.Generator {
	return new(generator)
}

func (g *generator) Train(s string) {}

func (g *generator) Generate(maxLength int) string {
	s := "hello, world"
	if len(s) > maxLength {
		return s[:maxLength]
	}
	return s
}
