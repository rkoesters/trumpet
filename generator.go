package main

// Generator gets trained by calls to Train and creates strings of text
// by calls to Generate.
type Generator interface {
	Train(s string)
	Generate(maxLength uint) string
}

type dummyGenerator struct{}

func (d *dummyGenerator) Train(s string) {}

func (d *dummyGenerator) Generate(maxLength uint) string {
	return "hello, world"
}
