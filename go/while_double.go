// The "While" Sum: Create a loop that acts like a while loop.
// Start with n = 1 and keep doubling it (n *= 2) until it is greater than 1,000.

package main

import "fmt"

func whileDouble() {
	// Short-hand declaration (idiomatic for local variables)
	n := 1

	for n <= 1_000 {
		fmt.Printf("Current value: %d\n", n)
		n *= 2
	}

	fmt.Printf("Final value: %d\n", n) // Output: 1024
}
