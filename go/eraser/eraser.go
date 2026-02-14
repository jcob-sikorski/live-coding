// The Eraser (Medium)
// Goal: Use the delete function and len.

// Create a map of int:string representing ID numbers and product names.

// Populate it with 5 items.

// Delete two items by their ID.

// Print the map and its final length.

package main

import "fmt"

func eraser() {
	products := map[int]string{
		101: "Milk",
		102: "Yoghurt",
		103: "Bread",
		104: "Eggs",
		105: "Cheese",
	}

	delete(products, 101)
	delete(products, 104)

	fmt.Println("Remaining Products:", products)
	fmt.Println("Final Length:", len(products))
}
