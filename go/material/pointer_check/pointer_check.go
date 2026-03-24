// The Basics (Warm-up)
// Exercise 1: The Identity Check
// Create an integer x. Create a pointer p that points to x.
// Print the memory address of x, the value of p,
// and the value stored at the address p is holding.

// Goal: Understand that a pointer is just a variable holding a memory address.

package main

import "fmt"

func checkIdentity() {
	// 1. Create an integer x and assign it a value
	var x int = 42

	// 2. Create a pointer p that points to x
	// The & operator "gets the address" of the variable
	var p *int = &x

	// 3. Print the results
	fmt.Println("Memory address of x:        ", &x)
	fmt.Println("Value of p (the address):   ", p)
	fmt.Println("Value at address p (*p):    ", *p)
}
