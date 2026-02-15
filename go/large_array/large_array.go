// The Large Array Challenge
// Create a function that declares a very large array (e.g., [1000000]int).
// Run it. Then, try making it even larger until you hit a stack overflow or
// see it automatically move to the heap.

// Goal: Learn that size is a factor in escape analysis.

package main

// This function takes a pointer but doesn't let it escape.
// go:noinline
func useArray(a *[10]int) {
	a[0] = 1
}

// go:noinline
func useLargeArray(a *[1000000]int) {
	a[0] = 1
}

func main() {
	// 1. Small Array - Should stay on stack now
	var small [10]int
	useArray(&small)

	// 2. Large Array - Will escape due to size (> 64KB)
	var large [1000000]int
	useLargeArray(&large)
}
