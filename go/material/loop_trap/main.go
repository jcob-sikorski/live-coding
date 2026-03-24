// The Loop Trap
// Start 10 goroutines inside a for loop. Each should print
// its loop index (e.g., "I am number 3").

// Challenge: Notice if the numbers are out of order or if
// they all print the same number (a common closure bug).

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i) // Problem: captures the variable 'i' by reference
		}()
	}
	wg.Wait()
}
