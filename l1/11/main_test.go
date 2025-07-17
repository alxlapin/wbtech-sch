package main

import (
	"slices"
	"testing"
)

func TestSlicesIntersection(t *testing.T) {
	tests := []struct {
		s1       []int
		s2       []int
		expected []int
	}{
		{
			[]int{3, 4, 5, 6},
			[]int{1, 2, 3, 4},
			[]int{3, 4},
		},
		{
			[]int{5, 6},
			[]int{1, 2, 3, 4},
			[]int{},
		},
		{
			[]int{1, 2, 3},
			[]int{1, 2, 3, 4},
			[]int{1, 2, 3},
		},
		{
			[]int{1, 2, 4, 8},
			[]int{9, 8, 5, 4},
			[]int{4, 8},
		},
		{
			[]int{},
			[]int{1, 2, 3, 4},
			[]int{},
		},
		{
			[]int{},
			[]int{},
			[]int{},
		},
	}

	for _, test := range tests {
		result := getSlicesIntersection(test.s1, test.s2)

		if !slices.Equal(test.expected, result) {
			t.Error("Expected", test.expected, "Got", result)
		}
	}
}
