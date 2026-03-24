// The Profile Updater
// Define a User struct with Name and Age. Write two methods:

// UpdateAgeValue(newAge int) (Value receiver)

// UpdateAgePointer(newAge int) (Pointer receiver)
// Call both in main and explain why one works and the other doesn't.

package main

type UserData struct {
	Name string
	Age  int
}

// UpdateAgePointer (Pointer receiver)
// This works on the memory address of the original struct.
func (u *UserData) UpdateAgePointer(newAge int) {
	u.Age = newAge
}

// UpdateAgeValue (Value receiver)
// This works on a copy of the struct.
func (u UserData) UpdateAgeValue(newAge int) {
	u.Age = newAge
}

// func main() {
// 	user := UserData{Name: "Alice", Age: 25}

// 	fmt.Printf("Original: %v\n", user)

// 	// 1. Trying Value Receiver
// 	user.UpdateAgeValue(30)
// 	fmt.Printf("After UpdateAgeValue: %v (No change!)\n", user)

// 	// 2. Trying Pointer Receiver
// 	user.UpdateAgePointer(30)
// 	fmt.Printf("After UpdateAgePointer: %v (Success!)\n", user)
// }
