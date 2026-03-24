package main

import "fmt"

func main() {
	x := 10
	defer fmt.Println("Value A:", x)

	x = 20
	defer func(val int) {
		fmt.Println("Value B:", val)
	}(x)

	x = 30
	fmt.Println("Value C:", x)
}
