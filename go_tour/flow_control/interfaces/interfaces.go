package main

import (
	"fmt"
	"math"
	"golang.org/x/tour/reader"

)
type MyFloat float64

type Abser interface {
	Abs() float64
}
type Vertex struct {
	X, Y float64
}
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if (t == nil){
		fmt.Println("nil from func M")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}



func (rd MyReader) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = 'A'
	}
	return len(p), nil
} 

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v

	fmt.Println(a.Abs())

	// Interface values with nil underlying values
	var i I
	describe(i) // (<nil>, <nil>)
	var t *T
	i = t
	describe(i) // (<nil>, *main.T)
	i = &T{"Hello"}
	describe(i) // (&{Hello}, *main.T)
	i.M() // Hello

	i = F(math.Pi) // describe(i) // (3.141592653589793, main.F)
	describe(i) // (3.141592653589793, main.F)
	i.M() // 3.141592653589793


	var j interface{}
	describe(j) // (<nil>, <nil>)

	j = 42
	describe(j) // (42, int)

	j = "hello"
	describe(j) // (hello, string)


	var inter interface{} = "hello"
	fmt.Println("your i is ", i)
	// i.(float64) = 2.0
	
	s := inter.(string)
	fmt.Println(s)

	s, ok := inter.(string)
	fmt.Println(s, ok)
	i = 2.0
	f, ok := inter.(float64)
	fmt.Println(f, ok)

	f = inter.(float64)
	fmt.Println(f)
	
	// s = inter.(string) // panic
	// fmt.Println(s)

}


