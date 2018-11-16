package count

import (
	"testing"
)

func TestCount(t *testing.T) {
	g := New()

	g.Train("test")
	g.Train("test")

	if uint64(*g) != 2 {
		t.Fail()
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
