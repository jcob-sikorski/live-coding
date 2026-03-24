// Define a Book struct with fields for Title (string), Author (string),
// and Pages (int). Create an instance in main using field names and
// print it using %+v to see the labels.

package main

import "fmt"

// Book defines the structure for our library entries
type Book struct {
	Title  string
	Author string
	Pages  int
}

func printBook() {
	// Create an instance using field names for all fields
	book := Book{
		Title:  "The Go Programming Language",
		Author: "Alan A. A. Donovan",
		Pages:  380,
	}

	// %+v prints the field names along with the values
	fmt.Printf("%+v\n", book)
}
