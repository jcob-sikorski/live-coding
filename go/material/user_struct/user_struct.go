// The Zero-Value Audit
// Create a User struct with Username, Email, Age, and IsAdmin.
// Declare it using var u User without assigning any values. Print it.
// Observe what happens to the strings, the int, and the boolean.

package main

import "fmt"

type User struct {
	Username string
	Email    string
	Age      int
	IsAdmin  bool
}

func declareUser() {
	var u User

	fmt.Println(u) // Username: "", Email: "", Age: 0, isAdmin: false
}
