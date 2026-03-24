// The "Constructor" Pattern
// Go doesn't have classes, but we use functions. Create a function NewUser(username, email string) *User.
// It should return a pointer to a new User with IsAdmin set to false by default.

package main

import "fmt"

type UserConfig struct {
	Username string
	Email    string
	IsAdmin  bool
}

func NewUserConfig(username string, email string) *UserConfig {
	return &UserConfig{
		Username: username,
		Email:    email,
		IsAdmin:  false, // Explicitly setting it for clarity
	}
}

func main() {
	// Usage
	user := NewUserConfig("gopher123", "gopher@golang.org")

	fmt.Printf("User: %s, Email: %s, Admin: %v\n", user.Username, user.Email, user.IsAdmin)
}
