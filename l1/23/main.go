package main

import (
	"fmt"
	"slices"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	i := 5

	copy(s[i:], s[i+1:])

	s = s[:len(s)-1]
	s = slices.Clip(s)

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
