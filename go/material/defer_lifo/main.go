// The "Reverse" Warm-up (Level: Easy)
// The Goal: Practice the LIFO (Last-In, First-Out) nature of defer.

// Task: Write a function that takes a string (e.g., "GOLANG") and
// uses defer within a loop to print the characters in reverse order ("GNALOG").

// Hint: Each character should be printed in its own deferred call.

package main

import "fmt"

func deferPrint(s string) {
	fmt.Printf("Original: %s\nReversed: ", s)

	for _, c := range s {
		// We convert the rune 'c' to a string so it prints as a letter
		defer fmt.Print(string(c))
	}
}

func main() {
	deferPrint("GOLANG")
	// Output:
	// Original: GOLANG
	// Reversed: GNALOG
}
