package main

import "fmt"

func quickSort(s []int) {
	if len(s) <= 1 {
		return
	}

	pivotIndex := len(s) - 1

	var i int
	for i = 0; i < pivotIndex; i++ {
		if s[i] > s[pivotIndex] {
			var j int
			for j = i + 1; j < pivotIndex; j++ {
				if s[j] <= s[pivotIndex] {
					s[i], s[j] = s[j], s[i]
					break
				}
			}

			if j == pivotIndex {
				s[i], s[pivotIndex] = s[pivotIndex], s[i]
				break
			}
		}
	}

	quickSort(s[:i])
	quickSort(s[i+1:])
}

func main() {
	slice := []int{4, 6, 8, 3, 5, 7, 9, 2}

	quickSort(slice)

	fmt.Println("Output =", slice)
}
