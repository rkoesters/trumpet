package twitter

import (
	"testing"
)

func TestIsStringInSlice(t *testing.T) {
	slice := []string{
		"test",
		"in slice",
		"test",
	}

	str1 := "in slice"
	if !isStringInSlice(str1, slice) {
		t.Fail()
	}

	str2 := "not in slice"
	if isStringInSlice(str2, slice) {
		t.Fail()
	}
}
