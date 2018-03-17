package count

type Generator uint64

func New() *Generator {
	return new(Generator)
}

func (g *Generator) Train(s string) {
	*g++
}

func (g *Generator) Generate(maxLength int) string {
	panic("count.Generate")
}
