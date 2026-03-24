// The Explicit Converter
// Declare var a int32 = 100 and var b int64 = 200.
// Try to add them and store the result in a variable c.

package main

import "fmt"

func main() {
	var a int32 = 100
	var b int64 = 200

	// You converted 'a' to match 'b'. This is the 'Go way'.
	var c int64 = int64(a) + b

	fmt.Printf("The result is: %d\n", c)
	fmt.Printf("Type of c: %T\n", c)
}
