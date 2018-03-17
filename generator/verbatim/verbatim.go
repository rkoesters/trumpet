package verbatim

type Generator struct {
	m map[string]struct{}
}

func New() *Generator {
	return &Generator{make(map[string]struct{})}
}

func (g *Generator) Train(s string) {
	var v struct{}
	g.m[s] = v
}

func (g *Generator) Generate(maxLength int) string {
	panic("verbatim can't generate")
}

func (g *Generator) Exists(s string) bool {
	_, ok := g.m[s]
	return ok
}
