package main

import (
	"fmt"
	"slices"
)

func getSlicesIntersection(s1 []int, s2 []int) []int {
	var intersection []int

	slices.Sort(s1)
	slices.Sort(s2)

	minSliceLength := min(len(s1), len(s2))

	i1 := 0
	i2 := 0
	for i1 < minSliceLength && i2 < minSliceLength {
		if s1[i1] == s2[i2] {
			intersection = append(intersection, s1[i1])
			i1++
			i2++
		} else if s1[i1] < s2[i2] {
			i1++
		} else {
			i2++
		}
	}

	return intersection
}

func main() {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{3, 4, 5, 6}

	fmt.Println("Output", getSlicesIntersection(s1, s2))
}
