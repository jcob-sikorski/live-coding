// The "Remote Control"
// Write a function double(n *int). It should take a pointer to an integer and
// multiply the original value by 2. Call it from main and verify that the original variable changed.

// Goal: Master the * dereferencing operator to modify data remotely.

package main

// double takes a pointer to an int and multiplies the value at that address by 2
func double(n *int) {
	// *n accesses the value stored at the memory address
	*n *= 2
}

// func main() {
// 	original := 21

// 	fmt.Printf("Before: %d (Address: %p)\n", original, &original)

// 	// We pass the memory address of 'original' using the & operator
// 	double(&original)

// 	fmt.Printf("After:  %d (Address: %p)\n", original, &original)

// 	if original == 42 {
// 		fmt.Println("Success! The original variable was modified.")
// 	}
// }
