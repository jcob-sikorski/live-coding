// The Pointer Return (Escape 101)
// Create two functions: one that returns an int and one that returns an *int.
// Use the -gcflags="-m" flag to see which one "escapes to heap."

// Goal: Observe how returning a reference forces data onto the heap.

package main

func returnValue() int {
	var x int
	return x
}

func returnPointer() *int {
	var x int
	return &x
}

func main() {
	returnValue()
	returnPointer()
}
