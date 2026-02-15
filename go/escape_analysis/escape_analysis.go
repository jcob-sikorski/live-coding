// The Escape Analysis Mystery
// Write a function that creates a local variable and returns a pointer to it.
// Does it crash? (No!) Research why this works in Go but fails in C.

package main

import "fmt"

func moveLocalToHeap() *int {
	var x int
	fmt.Println("x is on stack", x)
	return &x
}

func main() {
	y := moveLocalToHeap()
	fmt.Println("y is on heap", *y)
}
