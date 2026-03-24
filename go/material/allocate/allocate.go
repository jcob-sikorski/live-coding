// Pre-allocation (Efficiency)
// Goal: Use make.
// Create a slice of integers with a length of 0 but a capacity of 5. Use a loop to append 5 numbers. Print the length and capacity after each append to see how they change.

package main

import "fmt"

func allocate() {
	var nums []int = make([]int, 0, 5)

	for i := 0; i < 5; i++ {
		nums = append(nums, i)
		fmt.Println("Length:", len(nums))
		fmt.Println("Capacity:", cap(nums))
	}
}
