// The Interface Bridge
// Create two structs: Circle and Square. Give both an Area() float64 method.

// The Final Boss: Create an interface Shape that requires an Area() method.
// Write a function that takes a slice of Shape and prints the area of any
// struct passed to it.

package main

import (
	"fmt"
	"math"
)

// 1. Define the Structs
type Circle struct {
	Radius float64
}

type Square struct {
	Side float64
}

// 2. Implement the Area() methods
// Note: These methods now satisfy the Shape interface automatically.
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

// 3. The Final Boss: The Interface
// Any type that has an Area() method returning a float64 is now a 'Shape'.
type Shape interface {
	Area() float64
}

// 4. The Universal Function
func printAreas(shapes []Shape) {
	for _, shape := range shapes {
		fmt.Printf("Area of %T: %.2f\n", shape, shape.Area())
	}
}

func main() {
	shapes := []Shape{
		Circle{Radius: 5},
		Square{Side: 10},
		Circle{Radius: 2.5},
	}

	printAreas(shapes)
}
