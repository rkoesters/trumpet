// Package multi provides a trumpet.Generator that multiplexes training
// data over multiple trumpet.Generators.
package multi

import (
	"github.com/rkoesters/trumpet"
)

type Generator struct {
	trainers  []trumpet.Generator
	generator trumpet.Generator
}

func New() *Generator {
	return &Generator{}
}

func (g *Generator) AddTrainer(trainer trumpet.Generator) {
	g.trainers = append(g.trainers, trainer)
}

func (g *Generator) SetGenerator(generator trumpet.Generator) {
	g.generator = generator
}

func (g *Generator) Train(s string) {
	for _, trainer := range g.trainers {
		trainer.Train(s)
	}
}

func (g *Generator) Generate(maxLength int) string {
	return g.generator.Generate(maxLength)
}
