// The Safe Wrapper (Error Handling)
// Create a function ExecuteWithRecovery(task func()). It should run the task and use recover()
// inside a defer block to catch any panics, printing a friendly message instead of crashing.

package main

import (
	"fmt"
)

// ExecuteWithRecovery wraps a task to prevent application crashes.
func ExecuteWithRecovery(task func()) {
	defer func() {
		if r := recover(); r != nil {
			// Instead of just printing the raw panic value,
			// we add some context to make it "friendly."
			fmt.Printf("🚨 Recovery Alert: The task encountered an issue but was safely stopped.\n")
			fmt.Printf("Details: %v\n", r)
		}
	}()

	task()
}

func main() {
	fmt.Println("Starting the engine...")

	// Scenario 1: A task that works fine
	ExecuteWithRecovery(func() {
		fmt.Println("✅ Task 1: Successfully making coffee.")
	})

	// Scenario 2: A task that panics
	ExecuteWithRecovery(func() {
		fmt.Println("⚠️ Task 2: Trying to divide by zero...")
		// Manually triggering a panic for demonstration
		panic("Oh no, the coffee machine exploded!")
	})

	fmt.Println("🚀 Application is still running! We survived the panic.")
}
