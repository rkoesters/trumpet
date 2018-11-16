package markov

import (
	"testing"
)

func TestMarkov(t *testing.T) {
	const str = "one two three"

	g := NewChain(3)

	g.Train(str)

	if g.Generate(280) != str {
		t.Fail()
	}
}
