// The Basic Listener
// Create two goroutines that send different strings ("Ping" and "Pong")
// to two separate channels at different intervals. Use a select statement
// inside a loop to print whichever message arrives first.

package main

import (
	"fmt"
	"time"
)

func sender(ch chan<- string, message string, interval time.Duration) {
	for {
		time.Sleep(interval)
		ch <- message
	}
}

func main() {
	pingChan := make(chan string)
	pongChan := make(chan string)

	// Start the publishers with different intervals
	go sender(pingChan, "Ping", 500*time.Millisecond)
	go sender(pongChan, "Pong", 1200*time.Millisecond)

	fmt.Println("Listening for messages... (Ctrl+C to stop)")

	// The "Basic Listener" loop
	for {
		select {
		case msg1 := <-pingChan:
			fmt.Println("Received:", msg1)
		case msg2 := <-pongChan:
			fmt.Println("Received:", msg2)
		case <-time.After(2 * time.Second):
			// Safety timeout if nothing happens for 2 seconds
			fmt.Println("...still waiting...")
		}
	}
}
