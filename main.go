package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		channel1 <- "message to first channel"
	}()

	go func() {
		channel2 <- "message to second channel"
	}()

	// additional goroutines for each read operation
	// go func() {
	// 	for value := range channel1 {
	// 		fmt.Println(value)
	// 	}
	// }()

	// go func() {
	// 	for value := range channel2 {
	// 		fmt.Println(value)
	// 	}
	// }()

	// select statement
	timeout := time.After(5 * time.Second)
	go func() {
		for {
			select {
			case value := <-channel1:
				fmt.Println(value)
			case value := <-channel2:
				fmt.Println(value)
			case <-timeout:
				fmt.Println("5 seconds without messages")
				panic("no messages")
			default:
				time.Sleep(time.Second)
				fmt.Println("waiting...")
			}
		}
	}()

	fmt.Scanln()
}
