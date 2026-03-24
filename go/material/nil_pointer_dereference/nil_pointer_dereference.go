// The Nil Safety Check
// Create a function GetLength(s *string) int. If the pointer is nil, return 0.
// Otherwise, return the length of the string.
// Test it by passing both a valid pointer and a nil pointer.

// Goal: Learn to prevent the "nil pointer dereference" panic—the #1 Go bug.

package main

func GetLength(s *string) int {
	if s == nil {
		return 0
	}
	return len(*s)
}

// func main() {
//     // Test 1: A valid pointer
//     message := "Hello, Go!"
//     fmt.Printf("Valid pointer length: %d\n", GetLength(&message))

//     // Test 2: A nil pointer
//     var nilPtr *string
//     fmt.Printf("Nil pointer length: %d\n", GetLength(nilPtr))
// }
