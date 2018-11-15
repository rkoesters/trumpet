// Package verbatim provides a trumpet.Generator that keeps track of
// training data for the purposes of checking that generated data isn't
// identical to any piece of training data.
package verbatim

import (
	"strings"
	"sync"
)

// Generator is a trumpet.Generator that keeps track of training data.
// Exist can be called to check whether a string has been given to the
// Generator.
type Generator struct {
	m     map[string]struct{}
	mutex sync.Mutex
}

// New returns a *Generator.
func New() *Generator {
	return &Generator{
		m: make(map[string]struct{}),
	}
}

// Train adds the given string to the Generator's internal data
// structure of strings.
func (g *Generator) Train(s string) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	g.m[normalize(s)] = struct{}{}
}

// Generate panics.
func (g *Generator) Generate(maxLength int) string {
	panic("verbatim can't generate")
}

// Exists returns a bool indicating whether the given string has already
// been given to Train.
func (g *Generator) Exists(s string) bool {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	_, ok := g.m[s]
	return ok
}

func normalize(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
