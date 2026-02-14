// Dictionary Alphabetizer (Medium)
// Goal: Handle map's unordered nature.

// Create a map of words and definitions.

// Because maps are unordered, write code that prints the words and definitions in alphabetical order by the word (key).

// Hint: You’ll need a slice to help with this.

// dictionary := map[string]string{
// 	"Go":         "An open-source programming language.",
// 	"Algorithm":  "A process or set of rules to be followed in calculations.",
// 	"Boolean":    "A data type that has one of two possible values.",
// 	"Compiler":   "A program that converts instructions into a machine-code form.",
// 	"Data":       "Facts and statistics collected together for reference.",
// }

package main

import (
	"fmt"
	"sort"
)

func sortedMap(unsortedMap map[string]string) {
	dictionary := map[string]string{
		"Go":        "An open-source programming language.",
		"Algorithm": "A process or set of rules to be followed in calculations.",
		"Boolean":   "A data type that has one of two possible values.",
		"Compiler":  "A program that converts instructions into a machine-code form.",
		"Data":      "Facts and statistics collected together for reference.",
	}

	keys := make([]string, 0, len(dictionary))
	for key := range dictionary {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	fmt.Println("--- Alphabetical Dictionary ---")
	for _, word := range keys {
		fmt.Printf("%s: %s\n", word, dictionary[word])
	}
}
