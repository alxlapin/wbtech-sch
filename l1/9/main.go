package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	ints := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go func() {
		for i := range ch1 {
			ch2 <- int(math.Pow(float64(i), 2))

			time.Sleep(time.Second)
		}

		close(ch2)
	}()

	for _, v := range ints {
		ch1 <- v
	}

	close(ch1)

	for i := range ch2 {
		fmt.Println("Received", i, "from ch1")
	}
}
