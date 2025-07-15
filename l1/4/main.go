package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func worker(id int, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker ", id, "is shutting down...")

			time.Sleep(2 * time.Second)

			fmt.Println("Worker ", id, "stopped")
			return
		default:
			fmt.Printf("Worker %d doing work...\n", id)

			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT)

	defer stop()

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("Couldn't parse num workers")
	}

	var wg sync.WaitGroup

	wg.Add(numWorkers)

	for w := 1; w <= numWorkers; w++ {
		go worker(w, ctx, &wg)
	}

	fmt.Println("Waiting for the context to be done...")
	<-ctx.Done()
	fmt.Println("Received cancellation")

	wg.Wait()

	fmt.Println("Application stopped")
}
