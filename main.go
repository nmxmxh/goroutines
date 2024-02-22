package main

import (
	"fmt"
	"sync"
)

// solution two with mutex
func main() {
	var data int
	var mutex sync.Mutex

	go func() {
		mutex.Lock()
		defer mutex.Unlock()
		data = 64
	}()

	go func() {
		mutex.Lock()
		defer mutex.Unlock()
		fmt.Println(data)
	}()

	fmt.Scanln()
}

// solution one with atomic operations
// func main() {
// 	var data atomic.Int64

// 	go func() {
// 		data.Add(64)
// 	}()

// 	go func() {
// 		fmt.Println(data.Load())
// 	}()

// 	fmt.Scanln()
// }

// initial with data block
// func main() {
// 	var data int

// 	go func() {
// 		data = 42
// 	}()

// 	go func() {
// 		fmt.Println(data)
// 	}()

// 	fmt.Scanln()
// }
