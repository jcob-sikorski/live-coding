package main

import "fmt"

func callMe(count int) {
	// Each call adds a new "frame" to the stack
	if count%1000 == 0 {
		fmt.Printf("Stack depth: %d\n", count)
	}

	// Infinite recursion: The stack grows forever
	callMe(count + 1)
}

func main() {
	callMe(1)
}
