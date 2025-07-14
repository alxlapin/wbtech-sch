package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func worker(id int, bucket <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for j := range bucket {
		fmt.Println("worker", id, "received num", j)
	}
}

func main() {
	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("Couldn't parse num workers")
	}

	bucket := make(chan int, 100)

	var wg sync.WaitGroup

	wg.Add(numWorkers)

	for w := 1; w <= numWorkers; w++ {
		go worker(w, bucket, &wg)
	}

	for j := 1; j <= 100; j++ {
		bucket <- j
	}

	close(bucket)

	wg.Wait()
}
