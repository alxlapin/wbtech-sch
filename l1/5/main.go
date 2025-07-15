package main

import (
	"fmt"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала – читать эти значения. По истечении N секунд программа должна завершаться.
//
// Подсказка: используйте time.After или таймер для ограничения времени работы.
func main() {

	intBucket := make(chan int, 100)

	done := make(chan bool)

	timer := time.After(5 * time.Second)

	fmt.Println("Timer started at", time.Now())

	go func() {
		for {
			select {
			case <-timer:
				fmt.Println("Timer expired at", time.Now())

				done <- true
				return
			case v, ok := <-intBucket:
				if !ok {
					done <- true
					return
				}

				fmt.Println("Received", v)

				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	for i := 1; i <= 100; i++ {
		intBucket <- i
	}

	close(intBucket)

	<-done
}
