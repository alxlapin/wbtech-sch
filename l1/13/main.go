package main

import "fmt"

func swap(a int, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b

	return a, b
}

func main() {
	a := 1
	b := 2

	a, b = swap(a, b)

	fmt.Printf("a = %d b = %d\n", a, b)
}
