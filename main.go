package main

import (
	"fmt"
	"time"
)

func main() {
	// create an unbuffered channel.
	channel := make(chan string)

	// sender goroutine function.
	go func() {
		time.Sleep(5 * time.Second)
		channel <- "message"
	}()

	// receiver goroutine function.
	go func() {
		message := <-channel
		fmt.Printf("message: %v", message)
	}()

	fmt.Scanln()
}
