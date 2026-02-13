// The Countdown: Write a standard for loop that counts down from 10 to 1 and then prints "Blast off!"

package main

import "fmt"

func forLoop() {
	for x := 10; x > 0; x-- {
		fmt.Println(x)
	}
	fmt.Println("Blast off!")
}
