// Goal: Initialize a map and perform a basic lookup.

// Create a map where the keys are names (strings) and values are ages (ints).

// Add three people to the map.

// Print the age of one specific person using their name.

package main

func nameTag(names []string, ages []int) map[string]int {
	// 1. Find the SHORTEST length to avoid index out of range errors
	length := len(names)
	if len(ages) < length {
		length = len(ages)
	}

	// 2. Initialize the map
	tagMap := make(map[string]int)

	// 3. Zip the slices into the map
	for i := 0; i < length; i++ {
		tagMap[names[i]] = ages[i]
	}

	return tagMap
}
