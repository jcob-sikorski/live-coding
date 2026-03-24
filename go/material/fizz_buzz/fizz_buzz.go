// FizzBuzz (The Classic): Loop from 1 to 50. If divisible by 3, print "Fizz"; by 5, print "Buzz";
// by both, print "FizzBuzz."

package main

import "fmt"

func fizzBuzz() {
	for x := 1; x <= 50; x++ {
		switch {
		case x%15 == 0: // Check the most restrictive condition first
			fmt.Println("FizzBuzz")
		case x%3 == 0:
			fmt.Println("Fizz")
		case x%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(x)
		}
	}
}
