package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 1.
	// create a go routine
	// change a variable inside it.
	// ! incorrect value shown
	// goroutine didn't have time to finish before the program finishes.
	// add sleep
	var count int

	// 4.
	// let's solve that using mutex.
	var mu sync.Mutex

	// 2.
	// create new routines, loop.
	// ! different problem, random
	// not a problem of waiting.
	for i := 0; i < 1000; i++ {
		go func() {
			// 3.
			// print inside the goroutine.
			// ! a lot of goroutines changing variable at the same time.

			// 5.
			// lock and unlock in the function.
			// we will have order now.
			mu.Lock()
			defer mu.Unlock()
			fmt.Println(count)
			count++
		}()
	}

	time.Sleep(time.Second)
	fmt.Println(count)
}
