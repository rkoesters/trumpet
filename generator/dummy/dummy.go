// Package dummy provides a dummy trumpet.Generator.
package dummy

import (
	"log"
)

type Generator struct{}

func (g *Generator) Train(s string) {
	log.Printf("dummy: learned about %v", s)
}

func (g *Generator) Generate(maxLength int) string {
	s := "hello, world"
	if len(s) > maxLength {
		return s[:maxLength]
	}
	return s
}
