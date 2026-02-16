// The Signal to Stop
// Start a goroutine that prints "Working..." every 500ms in an infinite loop.
// Use a second channel called quit (type chan bool) to tell that goroutine
// to stop and exit when the user presses "Enter."

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(quit <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-quit:
			// If we receive any signal on the quit channel, exit the loop
			fmt.Println("Worker received stop signal!")
			return
		default:
			// Otherwise, do the work
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	quit := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(1)
	go worker(quit, &wg)

	fmt.Println("Worker started. Press 'Enter' to stop it...")

	// This blocks until the user presses Enter
	fmt.Scanln()

	// Send the signal to stop
	quit <- true

	// Wait for the worker to finish its cleanup
	wg.Wait()
	fmt.Println("Program exited cleanly.")
}
