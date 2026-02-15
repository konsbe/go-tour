package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	// c → data channel (numbers)
	c := make(chan int)
	// quit → stop signal channel
	quit := make(chan int)
	// the above go routine will run concurrently with the main goroutine, 
	// and it will send Fibonacci numbers to channel c 
	// until it receives a signal on channel quit.
	// through the select statement, the goroutine can wait on multiple communication operations.
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c, i)
			// it quits if i == 2
			// if i == 2 {
			// 	quit <- 0
			// }
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
