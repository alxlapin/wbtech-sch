package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func cancelGoroutineWhenConditionMet() {
	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		defer wg.Done()

		for i := 1; i < 10; i++ {
			fmt.Println("Горутина обрабатывает элемент №", i)

			time.Sleep(1 * time.Second)

			if i == 3 {
				fmt.Println("Сработало условие на выход из горутины")
				return
			}
		}
	}()

	wg.Wait()
}

func cancelGoroutineWithChannel() {
	done := make(chan int)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Канал закрыт. Горутина завершена")

				time.Sleep(500 * time.Millisecond)
				return
			default:
				fmt.Println("Горутина работает...")

				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	close(done)
	time.Sleep(2 * time.Second)
}

func cancelGoroutineWithContext() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Контекст закрыт. Горутина завершена")

				time.Sleep(500 * time.Millisecond)
				return
			default:
				fmt.Println("Горутина работает...")

				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}

func cancelGoroutineWithRuntimeExit() {
	// Ключевые моменты вызова [[runtime.Exit()]]:
	// 1. defer выполняется перед выходом
	// 2. корректное [неаварийное] завершение горутины
	// 3. останавливает только горутину, в которой вызвана
	go func() {
		defer fmt.Println("Сработал defer после вызова runtime.Goexit")

		fmt.Println("Горутина работает...")

		time.Sleep(1 * time.Second)

		runtime.Goexit()

		fmt.Println("Этот участок кода не выполнится, т.к. горутина завершается немедленно")
	}()

	time.Sleep(2 * time.Second)
}

func main() {
	fmt.Println("\n-----example cancelGoroutineWhenConditionMet")
	cancelGoroutineWhenConditionMet()
	fmt.Println("\n-----example cancelGoroutineWithChannel")
	cancelGoroutineWithChannel()
	fmt.Println("\n-----example cancelGoroutineWithContext")
	cancelGoroutineWithContext()
	fmt.Println("\n-----example cancelGoroutineWithRuntimeExit")
	cancelGoroutineWithRuntimeExit()
}
