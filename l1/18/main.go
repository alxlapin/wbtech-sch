package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	c     int
	mutex sync.Mutex
}

func (c *SafeCounter) Inc() {
	c.mutex.Lock()
	c.c++
	c.mutex.Unlock()
}

func main() {
	sc := SafeCounter{c: 0}

	wg := sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sc.Inc()
		}()
	}

	wg.Wait()

	fmt.Println("Output", sc.c)
}
