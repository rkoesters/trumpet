// Package verbatim provides a trumpet.Generator that keeps track of
// training data for the purposes of checking that generated data isn't
// identical to any piece of training data.
package verbatim

import (
	"strings"
	"sync"
)

type Generator struct {
	m     map[string]struct{}
	mutex sync.Mutex
}

func New() *Generator {
	return &Generator{
		m: make(map[string]struct{}),
	}
}

func (g *Generator) Train(s string) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	g.m[normalize(s)] = struct{}{}
}

func (g *Generator) Generate(maxLength int) string {
	panic("verbatim can't generate")
}

func (g *Generator) Exists(s string) bool {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	_, ok := g.m[s]
	return ok
}

func normalize(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
