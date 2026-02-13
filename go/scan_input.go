// The "Who Are You?" (Variables & Input)
// Write a program that asks for the user's name and then prints a personalized greeting. This will help you learn how to handle user input.

// Hint: Use fmt.Scanln() to capture what the user types.

// Bonus: Try asking for their age and print how many years they have until they turn 100.

package main

import "fmt"

func scanInput() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Hello,", name)
}
