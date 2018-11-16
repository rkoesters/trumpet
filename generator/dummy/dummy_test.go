package dummy

import (
	"testing"
)

func TestDummy(t *testing.T) {
	g := New()

	g.Train("test")

	if g.Generate(280) != "hello, world" {
		t.Fail()
	}

	if g.Generate(5) != "hello" {
		t.Fail()
	}
}
