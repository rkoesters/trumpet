package main

type Generator interface {
	Train(s string)
	Generate(maxLength uint) string
}

type dummyGenerator struct{}

func (d *dummyGenerator) Train(s string) {}

func (d *dummyGenerator) Generate(maxLength uint) string {
	return "hello, world"
}
