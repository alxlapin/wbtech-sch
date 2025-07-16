package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex

	m := make(map[string]int)

	for i := 0; i < 1000; i++ {
		go func() {
			// При отсутствии мьютекса результат непредсказуемый: зачастую - ошибка "concurrent map writes".
			mutex.Lock()
			m["count"]++
			mutex.Unlock()
		}()
	}

	time.Sleep(2 * time.Second)

	fmt.Println(m["count"])
}
