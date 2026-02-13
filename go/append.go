// The Grow Operation
// Goal: Use the append function.
// Start with an empty slice of integers. Use a for loop to append the numbers 1 through 10 to the slice. Print the final slice.

package main

import "fmt"

func appendInts() {
	var ints []int

	for x := 1; x <= 10; x++ {
		ints = append(ints, x)
	}

	fmt.Println("Final slice:", ints)
}
