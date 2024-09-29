package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3)

	routine := func() {
		fmt.Println("Goroutine: Waiting for a value from the channel...")
		fmt.Printf("Goroutine: Got the value %d from the channel.\n", <-ch)
	}

	go routine()
	go routine()
	go routine()
	// Populating the buffered channel to its limit.
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("Buffer is now full.")
	ch <- 4
	ch <- 5
	ch <- 6
	fmt.Println("This line prints after the 4th value gets through.")
}
