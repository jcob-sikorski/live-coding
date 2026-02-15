// The Value Copy
// Write a function double(n int) that takes an integer and tries to double it.
// In main, print the variable before and after calling the function. Observe why the original doesn't change.

// Goal: Understand that stack-allocated variables are passed by value (copied).

package main

import "fmt"

func double(n int) int {
	fmt.Printf("Address inside double: %p\n", &n)
	return 2 * n
}

func main() {
	n := 10
	fmt.Printf("Address inside main:   %p\n", &n)

	fmt.Println("n before doubling:", n)

	// The result of double(n) is returned, but 'n' here remains 10
	double(n)

	fmt.Println("n after doubling: ", n)
}
