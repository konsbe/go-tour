package main

import "fmt"
// Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
// (The data flows in the direction of the arrow.)
// By default, sends and receives block until the other side is ready. 
// This allows goroutines to synchronize without explicit locks or condition variables.
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func sum(s []int, c chan int) {
	sum := 0
	describe(s)
	// describe(c)
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
