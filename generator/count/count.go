// Package count provides a trumpet.Generator that keeps count of the
// number of times Train is called.
package count

// Generator is a trumpet.Generator that keeps track of the number of
// times that Train is called.
type Generator uint64

// New returns a *Generator that keeps track of the number of times
// Train is called.
func New() *Generator {
	return new(Generator)
}

// Train increments its counter.
func (g *Generator) Train(s string) {
	*g++
}

// Generate panics.
func (g *Generator) Generate(maxLength int) string {
	panic("count can't generate")
}
