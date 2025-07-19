package main

import "fmt"

func binarySearch(haystack []int, needle int) int {
	sLength := len(haystack)
	if sLength == 0 {
		return -1
	}

	i := sLength / 2
	if haystack[i] == needle {
		return i
	} else if needle < haystack[i] {
		return binarySearch(haystack[:i], needle)
	} else {
		res := binarySearch(haystack[i+1:], needle)
		if res != -1 {
			res = i + 1 + res
		}

		return res
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("Int = %d found at %d index\n", 2, binarySearch(s, 2))
	fmt.Printf("Int = %d found at %d index\n", 8, binarySearch(s, 8))
	fmt.Printf("Int = %d found at %d index\n", 0, binarySearch(s, 0))
	fmt.Printf("Int = %d found at %d index\n", 10, binarySearch(s, 10))
}
