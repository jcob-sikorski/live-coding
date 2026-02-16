// The Ping-Pong
// Create a program where the main goroutine sends the string "ping" to a channel,
// and a separate goroutine receives it and prints it.

package main

import (
	"fmt"
	"sync"
)

// receiver takes a receive-only channel of strings
func receiver(ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Receive the message from the channel
	msg := <-ch
	fmt.Println("Received message:", msg)
}

func main() {
	var wg sync.WaitGroup

	// Create a channel for strings
	ch := make(chan string)

	wg.Add(1)

	// Start the separate goroutine
	go receiver(ch, &wg)

	// Main goroutine sends "ping" to the channel
	fmt.Println("Sending: ping")
	ch <- "ping"

	// Wait for the receiver to finish
	wg.Wait()
}
