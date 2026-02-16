// Randomized Fairness
// Create two channels and fill them both with 10 messages immediately.
// Run a loop with a select statement 10 times to read from them.
// Observe and explain why the output isn't 5 messages from one and then 5 from the other.

package main

import (
	"fmt"
)

func main() {
	// 1. Create buffered channels with capacity for 10 messages each
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	// 2. Fill them both immediately
	for i := 1; i <= 10; i++ {
		ch1 <- i
		ch2 <- i
	}

	// 3. Run a loop 10 times to read from them
	fmt.Println("Starting consumption...")
	for i := 1; i <= 10; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("Iteration %d: Received %d from ch1\n", i, msg1)
		case msg2 := <-ch2:
			fmt.Printf("Iteration %d: Received %d from ch2\n", i, msg2)
		}
	}
}
