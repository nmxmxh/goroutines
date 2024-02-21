package main

import "fmt"

func main() {
	// add size (1) to buffer.
	channel := make(chan string, 1)

	// send a message to channel.
	// blocking thread.
	// in this case, main thread.
	channel <- "message"

	// ? add another message
	// ! channel is full
	channel <- "another message"

	fmt.Println(<-channel)
}
