package trumpet

// Generator gets trained by calls to Train and creates strings of text
// by calls to Generate.
type Generator interface {
	Train(s string)
	Generate(maxLength int) string
}
