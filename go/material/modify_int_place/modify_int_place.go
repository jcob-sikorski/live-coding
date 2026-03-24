// The Filter Challenge
// Goal: Manipulate slices in-place.
// Write a function FilterEven(nums []int) []int that takes a slice and returns
// a new slice containing only the even numbers.
// Bonus: Try to do this without allocating a second underlying array (modify the slice in-place).

package main

func FilterEven(nums []int) []int {
	n := nums[:0]

	for _, x := range nums {
		if x%2 == 0 {
			n = append(n, x)
		}
	}

	return n
}
