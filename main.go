package main

import (
	"fmt"
	"time"
)

func main() {
	go printMessage("print routine 1")
	go printMessage("print routine 2")
	go printMessage("print routine 3")
	go printMessage("print routine 4")

	// 1.
	// main goroutine. MAIN THREAD.
	// main thread initially not allowing others to execute.
	// added time sleep
	printMessage("print routine 5")

	// 4.
	// new function before sleep.
	// anonymous function synthax.
	// ! no control over order
	go func() {
		fmt.Println("print anonymous func routine")
	}()

	// 2.
	// time sleep
	// sleeping for one second
	// enough to execute other go routines.
	time.Sleep(time.Second)
}

func printMessage(msg string) {
	// 3.
	// iterating five times.
	for i := 0; i < 5; i++ {
		fmt.Println(msg)
	}

}
