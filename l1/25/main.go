package main

import (
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	timer := time.NewTimer(d)

	<-timer.C

	// или
	//<-time.After(d)
}

func main() {
	fmt.Println(time.Now())
	sleep(time.Second * 5)
	fmt.Println(time.Now())
}
