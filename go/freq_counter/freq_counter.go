// Word Frequency Counter (Easy)
// Goal: Use a map to aggregate data.

// Given a slice of strings: []string{"apple", "banana", "apple", "cherry", "banana", "apple"}.

// Create a map to count how many times each fruit appears.

// Print the final map.

package main

import "fmt"

func countFreq(fruits []string) map[string]int {
	counter := make(map[string]int)

	for _, fruit := range fruits {
		counter[fruit]++
	}

	fmt.Println(counter)

	return counter
}
