package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter = 0
	mutex   sync.Mutex
	cond    = sync.NewCond(&mutex)
)

func main() {
	go producer()
	go consumer()

	time.Sleep(5 * time.Second)
}

func producer() {
	for {
		mutex.Lock()
		if counter > 0 {
			cond.Wait()
		}

		time.Sleep(time.Second)
		counter++
		fmt.Printf("increasing counter: %v\n", counter)

		mutex.Unlock()
		cond.Signal()
	}
}

func consumer() {
	for {
		mutex.Lock()
		if counter == 0 {
			cond.Wait()
		}

		time.Sleep(time.Second)
		counter--
		fmt.Printf("decreasing counter: %v\n", counter)

		mutex.Unlock()
		cond.Signal()
	}
}
