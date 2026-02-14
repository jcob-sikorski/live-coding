// The Value Receiver
// Add a method Describe() to your Bible struct from Exercise 1.
// It should return a string like: " 'The Great Gatsby' by F. Scott Fitzgerald".
// Call it on a Bible instance.

package main

import "fmt"

// Bible defines the structure for our library entries
type Bible struct {
	Title  string
	Author string
	Pages  int
}

// Describe is a method with a value receiver (b Bible)
// It returns a formatted string about the Bible.
func (b Bible) Describe() string {
	return fmt.Sprintf("'%s' by %s", b.Title, b.Author)
}

func printBibleDescription() {
	// Create an instance of Bible
	myBible := Bible{
		Title:  "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Pages:  180,
	}

	// Call the method and print the result
	description := myBible.Describe()
	fmt.Println(description)
}
