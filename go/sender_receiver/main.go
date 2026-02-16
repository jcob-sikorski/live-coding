// The Directional Gatekeeper
// Write a function func sendOnly(ch chan<- int) that sends a value,and func receiveOnly(ch <-chan int).
// Pass the same channel to both from main to see how Go enforces "send-only" and "receive-only" constraints.

package main

import (
	"fmt"
	"sync"
)

func sendOnly(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch)
	ch <- 1
}

func receiveOnly(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(<-ch)
}

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(2)

	go sendOnly(ch, &wg)
	go receiveOnly(ch, &wg)

	wg.Wait()
}
