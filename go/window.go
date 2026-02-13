// The Window View
// Goal: Understand slicing syntax.
// Given the slice nums := []int{10, 20, 30, 40, 50}, create a new slice sub that contains only the middle elements [20, 30, 40]. Change the first element of sub to 99. Print the original nums slice.
// Hint: Observe if the original changed.

package main

import "fmt"

func window() {
	nums := []int{10, 20, 30, 40, 50}

	// Index 1 is '20', index 4 is '50'.
	// nums[1:4] gives us indices 1, 2, and 3.
	sub := nums[1:4]

	fmt.Println("Before change - sub:", sub)

	// Changing the first element of 'sub' (which is nums[1])
	sub[0] = 99

	fmt.Println("After change - sub:", sub)
	fmt.Println("Original nums:", nums) // Spoilers: nums[1] is now 99!
}
