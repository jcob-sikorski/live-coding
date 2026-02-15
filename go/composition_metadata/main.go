// Composition & Metadata (Advanced)
// 7. Struct Embedding (Composition)
// Create a Vehicle struct with Make and Model.
// Then, create a Car struct that embeds Vehicle and adds Doors.

// Goal: Initialize a Car and access myCar.Make directly (demonstrating field promotion).

package main

import "fmt"

type Vehicle struct {
	Make  string
	Model string
}

type Car struct {
	Vehicle
	Doors int
}

func main() {
	myCar := Car{
		Vehicle: Vehicle{
			Make:  "Toyota",
			Model: "Camry",
		},
		Doors: 4,
	}

	// Field Promotion:
	// Even though 'Make' is inside 'Vehicle', we can access it directly
	// from 'myCar' as if it were a top-level field.
	fmt.Println("Make:", myCar.Make)   // Accesses myCar.Vehicle.Make
	fmt.Println("Model:", myCar.Model) // Accesses myCar.Vehicle.Model
	fmt.Println("Doors:", myCar.Doors)
}
