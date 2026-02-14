// The Slice Trick
// Create a function ModifySlice(s []int). Inside, change s[0] = 999.
// In main, see if the original slice changed. Now, try to append a new number
// to the slice inside the function. Does the original slice in main show the new number?
// (Hint: It won't!)

package main

import "fmt"

func ModifySlice(s []int) {
	// 1. This affects the underlying array
	s[0] = 999

	// 2. This modifies the LOCAL copy of the slice header
	s = append(s, 888)
	fmt.Printf("Inside function: %v (len: %d, cap: %d)\n", s, len(s), cap(s))
}

// func main() {
// 	s := []int{1, 2, 3}

// 	fmt.Printf("Before main:    %v (len: %d, cap: %d)\n", s, len(s), cap(s))

// 	ModifySlice(s)

// 	fmt.Printf("After main:     %v (len: %d, cap: %d)\n", s, len(s), cap(s))
// }
