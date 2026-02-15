package main

import "fmt"

func main() {
	// 	make(chan int, 2) creates a channel that can hold int values
	// The 2 is the buffer capacity - it can store up to 2 values without blocking
	ch := make(chan int, 2) // Buffered channel with a capacity of 2
	ch <- 1  // Doesn't block (buffer has space)
	ch <- 2  // Doesn't block (buffer still has space)
	// ch <- 3 would block here (buffer full, no receiver)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
