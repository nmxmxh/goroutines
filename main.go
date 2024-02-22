package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter   = 0
	mutex     sync.Mutex
	condition = sync.NewCond(&mutex)
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
			condition.Wait()
		}
		time.Sleep(1 * time.Second)
		counter++
		fmt.Printf("incrementing %v \n", counter)
		mutex.Unlock()
		condition.Signal()
	}
}

func consumer() {
	for {
		mutex.Lock()
		if counter == 0 {
			condition.Wait()
		}
		time.Sleep(1 * time.Second)
		counter--
		fmt.Printf("decrementing %v \n", counter)
		mutex.Unlock()
		condition.Signal()
	}
}
