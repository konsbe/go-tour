package main

import (
	"fmt"
	"golang.org/x/example/hello/reverse"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(reverse.String("Hello, World!"))
	fmt.Println(reverse.String("Hello"), reverse.Int(24601))

}
