// The "Quit" Signal
// Create a worker goroutine that performs a task in an infinite loop.
// Use select to listen for either a "task" channel or a "quit" channel.
// If the quit channel receives a signal, the worker should print "Cleaning up..." and return.

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(tasks <-chan int, quit <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case task := <-tasks:
			fmt.Printf("Processing task %d\n", task)
		case <-quit:
			fmt.Println("Cleaning up...")
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	// Using struct{} saves memory; it's the standard for signals
	quit := make(chan struct{})
	tasks := make(chan int)

	wg.Add(1)
	go worker(tasks, quit, &wg)

	// Send some work
	for i := 1; i <= 3; i++ {
		tasks <- i
		time.Sleep(200 * time.Millisecond)
	}

	// Signal the worker to stop
	close(quit)

	wg.Wait()
	fmt.Println("Main exited.")
}
