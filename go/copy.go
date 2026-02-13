// The "Copy" Cat
// Goal: Prevent side effects using copy.
// You have a slice original := []int{1, 2, 3}. Create a new slice duplicate that has the same elements
// but is not connected to the same underlying array.
// If you change duplicate[0], original[0] should remain 1.

package main

import "fmt"

func duplicate() {
	var original []int = []int{1, 2, 3}

	var duplicated []int = make([]int, len(original), 10)

	copy(duplicated, original)

	original = append(original, 100)

	fmt.Println("Original:", original)
	fmt.Println("Duplicated:", duplicated)
}
