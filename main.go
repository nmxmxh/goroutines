package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var mu sync.Mutex
	var waitGroup sync.WaitGroup

	itrs := 100000
	// waitGroup.Add(itrs)

	// option 2. add individually.
	waitGroup.Add(1)

	go func() {
		// option 2. defer waitgroup done @ beginning.
		defer waitGroup.Done()
		for i := 0; i < itrs; i++ {
			waitGroup.Add(1)
			go func() {
				// option 1.
				defer waitGroup.Done()
				mu.Lock()
				count++
				mu.Unlock()

			}()
		}
		fmt.Print("first routine done.\n")
	}()

	altItrs := 1000
	// option 1.
	// waitGroup.Add(altItrs)

	// option 2. add individually.
	waitGroup.Add(1)

	go func() {
		defer waitGroup.Done()
		for i := 0; i < altItrs; i++ {
			waitGroup.Add(1)
			go func() {
				// option 2. put at beginning.
				defer waitGroup.Done()
				mu.Lock()
				count++
				mu.Unlock()
				// option 1.
				// waitGroup.Done()
			}()
		}
		fmt.Print("second routine done.\n")
	}()

	waitGroup.Wait()
	fmt.Println(count)
}
