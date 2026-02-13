// Type Switch (Advanced): In Go, you can switch on a variable's type.
// Create an interface{} variable and write a switch to detect
// if it’s a string, int, or bool.

package main

import "fmt"

func explainType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %q (Length: %d)\n", v, len(v))
	case bool:
		fmt.Printf("Boolean: %t\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}
