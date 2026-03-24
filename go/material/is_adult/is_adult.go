// The Bouncer: Write a script that checks if a variable age is 21 or older. Print "Enter" or "Access Denied."

package main

import "fmt"

func isAdult(age int) {
	if age >= 21 {
		fmt.Println("Enter")
	} else {
		fmt.Println("Access Denied")
	}
}

// if age := getAge(); age >= 21 {
// 	fmt.Println("Enter")
// } else {
// 	fmt.Println("Access Denied")
// }
