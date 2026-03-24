// Memory-Safe MathCreate a function that takes two int16 values, multiplies them,
// and returns an int32.Goal: Learn to "upcast" before an operation to prevent
// overflow (since $30,000 \times 30,000$ won't fit in an int16).

package main

import (
	"fmt"
	"math"
)

func upcastMultiply(a int16, b int16) int32 {
	return int32(a) * int32(b)
}

func main() {
	var val1 int16 = math.MaxInt16 // 32,767
	var val2 int16 = math.MaxInt16 // 32,767

	result := upcastMultiply(val1, val2)

	fmt.Printf("Input A: %d\n", val1)
	fmt.Printf("Input B: %d\n", val2)
	fmt.Printf("Result (int32): %d\n", result)
}
