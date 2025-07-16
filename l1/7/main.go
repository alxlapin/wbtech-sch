package main

import (
	"fmt"
	"sync"
)

func main() {
	var mutex sync.Mutex

	m := make(map[string]int)

	for i := 0; i < 1000; i++ {
		go func() {
			// При отсутствии мьютекса результат непредсказуемый: зачастую - ошибка "concurrent map writes"
			mutex.Lock()
			m["count"]++
			mutex.Unlock()
		}()
	}

	// Если убрать мьютекс при чтении, то запуск команды "go run -race" определит состояние гонки (см. пример в README)
	mutex.Lock()
	fmt.Println(m["count"])
	mutex.Unlock()
}
