// The ASCII/Byte Loop
// Write a loop that prints the numbers from 65 to 90. In the same line,
// convert that number to a string or byte to show the corresponding capital letter (A-Z).

// Goal: Understand the relationship between uint8, byte, and ASCII.

package main

import "fmt"

func main() {
	for x := 65; x <= 90; x++ {
		// %d formats the integer, %c formats the character (ASCII)
		// \t adds a tab for clean spacing
		fmt.Printf("%d \t %c\n", x, x)
	}
}
