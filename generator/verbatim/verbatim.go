// Package verbatim provides a trumpet.Generator that keeps track of
// training data for the purposes of checking that generated data isn't
// identical to any piece of training data.
package verbatim

type Generator struct {
	m map[string]struct{}
}

func New() *Generator {
	return &Generator{make(map[string]struct{})}
}

func (g *Generator) Train(s string) {
	g.m[s] = struct{}{}
}

func (g *Generator) Generate(maxLength int) string {
	panic("verbatim can't generate")
}

func (g *Generator) Exists(s string) bool {
	_, ok := g.m[s]
	return ok
}
