// The Buffered Buffer
// Create a buffered channel with a capacity of 3. Send 3 integers into it
// without starting any other goroutines. Then, receive and print them.

// Goal: Understand why buffered channels don't block immediately.

package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	for i := 0; i < 3; i++ {
		ch <- i
	}

	for i := 0; i < 3; i++ {
		fmt.Println("Value from channel", <-ch)
	}
}
