// The "Default" Explorer
// Create a program that declares an int, a float64, a byte,
// and a rune without assigning them values. Print them out.

package main

import "fmt"

func main() {
	var x int     // Default for integers
	var y float64 // Default for floats
	var z byte    // Alias for uint8
	var a rune    // Alias for int32

	fmt.Printf("int: %v\n", x)
	fmt.Printf("float64: %v\n", y)
	fmt.Printf("byte: %v\n", z)
	fmt.Printf("rune: %v\n", a)
}
