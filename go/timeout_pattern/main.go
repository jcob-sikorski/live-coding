// The Timeout Pattern
// Write a function that waits for a value on a channel.
// If the value doesn't arrive within 500 milliseconds,
// the select statement should print "Operation timed out" and exit.

package main

import (
	"fmt"
	"sync"
	"time"
)

func wait(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	select {
	case val := <-ch:
		fmt.Println("Received value", val)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("Operation timed out")
	}
}

func main() {
	wg := sync.WaitGroup{}

	ch := make(chan int)

	wg.Add(1)

	go wait(ch, &wg)

	ch <- 1

	wg.Wait()
}
