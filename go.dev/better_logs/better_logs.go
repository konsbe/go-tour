package main

import "fmt"

import "rsc.io/quote"

func main() {
	fmt.Println(quote.Go())
	// Go proverb about using channels for concurrency instead of shared memory with locks.
	// Don't communicate by sharing memory, share memory by communicating.
}
