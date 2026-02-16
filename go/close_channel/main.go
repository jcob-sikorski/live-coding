// The Range Loop
// Create a function that sends the numbers 1 to 10 into a channel and then closes it.
// In main, use a for range loop to print all the numbers.

// Goal: Practice closing channels to prevent deadlocks.

package main

import (
	"fmt"
)

// publisher only needs the channel
func publisher(ch chan int) {
	// defer close is the "gold standard" for preventing deadlocks
	defer close(ch)

	for i := 1; i <= 10; i++ {
		ch <- i
	}
}

func main() {
	ch := make(chan int, 10)

	// Fire off the goroutine
	go publisher(ch)

	// The range loop blocks main until the channel is closed.
	// This effectively "waits" for the publisher to finish.
	for val := range ch {
		fmt.Println(val)
	}

	fmt.Println("Done!")
}
