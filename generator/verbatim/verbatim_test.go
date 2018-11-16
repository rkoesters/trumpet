package verbatim

import (
	"testing"
)

func TestGenerator(t *testing.T) {
	in := []string{
		"",
		"test",
		"test 1",
		"test 2",
		"asdf",
	}

	not := []string{
		"test ",
		"test 3",
		"fdsa",
		" ",
	}

	g := New()

	for _, s := range in {
		g.Train(s)
	}

	for _, s := range in {
		if !g.Exists(s) {
			t.Fail()
		}
	}

	for _, s := range not {
		if g.Exists(s) {
			t.Fail()
		}
	}
}

func TestPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	g := New()

	_ = g.Generate(280)

	t.Fail()
}
