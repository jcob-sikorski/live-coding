// The "Ghost" Routine
// Create a function that prints "Inside Goroutine".
// Call it with the go keyword in main. Run it. Notice why it doesn't print anything,
// then add a time.Sleep to the main function to fix it temporarily.

package main

import (
	"fmt"
	"time"
)

func insideGoroutine() {
	fmt.Println("Inside Goroutine")
}

func main() {
	go insideGoroutine()
	time.Sleep(10000000)
}
