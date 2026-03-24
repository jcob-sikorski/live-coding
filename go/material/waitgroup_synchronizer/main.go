// The WaitGroup Synchronizer
// Refactor Exercise 1. Instead of using time.Sleep, use sync.WaitGroup
// to ensure the program waits for exactly one goroutine to finish before exiting.

package main

import (
	"fmt"
	"sync" // Import the sync package
)

func insideGoroutine(wg *sync.WaitGroup) {
	// Schedule the counter to decrement when the function returns
	defer wg.Done()
	fmt.Println("Inside Goroutine")
}

func main() {
	var wg sync.WaitGroup

	// Increment the counter because we are launching 1 goroutine
	wg.Add(1)

	// Pass a pointer to the WaitGroup so the function uses the same counter
	go insideGoroutine(&wg)

	// This blocks main until the counter is 0
	wg.Wait()

	fmt.Println("Main goroutine exiting gracefully.")
}
