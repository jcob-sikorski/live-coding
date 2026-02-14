// The Anonymous Quickie
// In your main function, create an anonymous struct to represent a point on a 2D plane (x and y).
// Initialize it and print the coordinates. This is useful for one-off data shapes!

package main

import "fmt"

func anonymousStruct() {
	// Defining and initializing the anonymous struct in one go
	point := struct {
		X float64
		Y float64
	}{
		X: 1.34,
		Y: 4.39,
	}

	fmt.Printf("Point Coordinates: (%.2f, %.2f)\n", point.X, point.Y)
}
