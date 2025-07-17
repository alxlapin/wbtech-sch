package main

import "fmt"

func main() {
	a := 1
	b := 2

	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("a = %d b = %d\n", a, b)
}
