// Package dummy provides a dummy trumpet.Generator.
package dummy

type Generator struct{}

func (g *Generator) Train(s string) {}

func (g *Generator) Generate(maxLength int) string {
	s := "hello, world"
	if len(s) > maxLength {
		return s[:maxLength]
	}
	return s
}
