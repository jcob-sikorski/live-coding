// The Weekender: Create a switch that takes an integer (1-7) and prints
// the name of the day. Use a default case for invalid numbers.

package main

import "fmt"

func basicSwitch(num int) {
	switch num {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6, 7: // You can comma-separate multiple values!
		fmt.Println("Weekend!")
	default:
		fmt.Println("Invalid day number. Please use 1-7.")
	}
}
