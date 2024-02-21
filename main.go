package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	go sender(channel)
	go receiver(channel)

	fmt.Scanln()
}

// chan<- in the function implies send only channel.
// ! if we try to read, receive an error
// assure it will only be used to send messages
func sender(channel chan<- string) {
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second)
		channel <- fmt.Sprintf("countdown for %v seconds", i)
	}
}

// chan<- in the function implies receive only channel.
// ! if we try to send, receive an error
// assure it will only be used to receive messages
func receiver(channel <-chan string) {
	// loop through using for statement.
	// with range as channel.
	for message := range channel {
		fmt.Println(message)
	}
}
