package main

import (
	"fmt"
	"math"
	"sync"
)

func pow(s []int, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, v := range s {
		fmt.Printf("Pow of %v: %v\n", v, math.Pow(float64(v), 2))
	}
}

func main() {
	inputIntArray := [5]int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	wg.Add(2)

	go pow(inputIntArray[:3], &wg)
	go pow(inputIntArray[3:], &wg)

	wg.Wait()

	fmt.Println("Goroutines are done")
}
