package main

import (
	"testing"
)

func TestSwap(t *testing.T) {
	tests := []struct {
		a int
		b int
	}{
		{1, 2},
		{0, 2},
		{-1, 2},
		{-1, -2},
	}

	for _, test := range tests {
		a, b := swap(test.a, test.b)

		t.Logf("a = %d b = %d\n", a, b)

		if a != test.b || b != test.a {
			t.Error("Something went wrong")
		}
	}
}
