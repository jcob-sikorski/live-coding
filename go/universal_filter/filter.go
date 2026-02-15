// The Universal Filter
// Write a function Filter(slice []int, test func(int) bool) []int.
// It should return a new slice containing only elements that pass the test.

package main

import "fmt"

func Filter[T any](a []T, test func(T) bool) []T {
	var b []T = make([]T, 0, len(a))
	for _, x := range a {
		if test(x) {
			b = append(b, x)
		}
	}
	return b
}

func main() {
	numbers := []int{1, 2, 3, 4, 11, 12, 13}

	// Filter for even numbers
	evens := Filter(numbers, func(n int) bool {
		return n%2 == 0
	})

	// Filter for numbers > 10
	largeNumbers := Filter(numbers, func(n int) bool {
		return n > 10
	})

	fmt.Println("Original:", numbers)             // [1 2 3 4 11 12 13]
	fmt.Println("Evens:", evens)                  // [2 4 12]
	fmt.Println("Greater than 10:", largeNumbers) // [11 12 13]
}
