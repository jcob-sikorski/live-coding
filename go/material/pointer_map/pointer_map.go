// The Map Mutation
// Create a map of string:int. Pass it to a function that adds a new key-value pair.
// Check the map in main. Do you need a pointer to a map (*map) for this to work?

// Goal: Recognize which types in Go carry their own pointers internally.

package main

import "fmt"

// modifyMap takes a map by value.
func modifyMap(m map[string]int) {
	m["added_key"] = 42
}

func main() {
	// Initializing the map
	m := make(map[string]int)
	m["original_key"] = 10

	fmt.Println("Before:", m)

	// Passing the map (by value) to the function
	modifyMap(m)

	// The map is modified in main!
	fmt.Println("After: ", m)
}
