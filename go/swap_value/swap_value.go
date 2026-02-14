// Create a function square(n *int) that takes a pointer to an integer and replaces the value at that address with its square.

// Goal: Understand how to dereference a pointer to modify a value.

package main

func square(n *int) {
	*n = (*n) * (*n)
}
