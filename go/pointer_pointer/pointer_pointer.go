// Advanced Concepts
// Exercise 7: The Pointer to Pointer
// Create an integer a. Create pointer p pointing to a. Create pointer pp pointing to p (**int). Use pp to change the value of a to 100.

// Goal: Understand multi-level indirection (rare in Go, but important for understanding how the engine works).

package main

import "fmt"

func main() {
	var a int
	var b *int = &a
	var bb **int = &b

	// Using the pointer to pointer to change 'a'
	**bb = 100

	fmt.Printf("Value of a: %d\n", a)
	fmt.Printf("Address of a: %p\n", &a)
	fmt.Printf("Value stored in b: %p\n", b)
	fmt.Printf("Value stored in bb: %p\n", bb)
}
