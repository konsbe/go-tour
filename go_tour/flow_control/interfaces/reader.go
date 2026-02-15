package main

import (
	"fmt"
	"math"
	"golang.org/x/tour/reader"

)

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}


func (rd MyReader) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = 'A'
	}
	return len(p), nil
} 

func main() {

	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z) // Arthur Dent (42 years) Zaphod Beeblebrox (9001 years)

	fmt.Println(Sqrt(2)) // 1.414213562373095 <nil>
	fmt.Println(Sqrt(-2)) // 0 cannot Sqrt negative number: -2

	reader.Validate(MyReader{})

}


