package main

import "fmt"

type Vertex struct {
	X int
	Y int
}
var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // allocates a new Vertex, has type *Vertex, takes its address, stores it in p
	// Vertex{...}     // value
	// *Vertex       → address of a value
	// &Vertex{...}    // pointer to a new value
	// v               // variable
	// &v              // pointer to existing variable
	// p               // pointer
	// &p              // pointer to pointer
)

func main() {
	i, j := 42, 2701

	p := &i         // & point to i
	fmt.Println(*p) // * access value through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i
	fmt.Println(&i) // 0xc000012050, the address of i
	fmt.Println(p) // 0xc000012050, the same address as i, because p points to i
	fmt.Println(&p) // 0xc00000e030, the address of p, which is different from i
	p = &j         // & point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j


	v := Vertex{1, 2}
	j := v // copy of v to j
	j.X = 8 // modify j.X
	p := &v // & point to v
	p.X = 2 // implicit dereferencing for struct fields, same as (*p).X = 2
	fmt.Println(v)
	fmt.Println(p)
	fmt.Println(j)

	fmt.Println(v1, p, v2, v3) // {1 2} &{1 2} {1 0} {0 0}
	// v1, v2, v3 → print values
	fmt.Println(&p) // 0x566f60 where the pointer lives
	fmt.Println(p) // &{1 2} → prints the value it points to, where the vertex lives
	fmt.Println(&v1) // &{1 2} → prints the address of v1, same as p


	fmt.Println(v1, p, v2, v3) // {1 2} &{1 2} {1 0} {0 0}
	// v1, v2, v3 → print values
	fmt.Println(&p) // 0x566f60 where the pointer lives
	fmt.Println(p) // &{1 2} → prints the value it points to, where the vertex lives
	fmt.Println(&v1) // &{1 2} → prints the address of v1, p != &v1   // always true
	k := &v1
	fmt.Println(*k) // {1 2}
	fmt.Println(&k) // 0xc00011c030


	// arrays, slices & methods
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	
	fmt.Println(a[0], a[1])
		names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b) 	// [John XXX] [XXX George]
	fmt.Println(names) 	// [John XXX George Ringo]


	//slices
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s) // [{2 true} {3 false} {5 true} {7 true} {11 false} {13 true}]
	// When slicing, you may omit the high or low bounds to use their defaults instead.
		
	ps := s[1:4]
	fmt.Println(ps) // [{3 false} {5 true} {7 true}]
	ps = s[:2]
	fmt.Println(ps) // [{2 true} {3 false}]
	ps = s[1:]
	fmt.Println(ps) // [{3 false} {5 true} {7 true} {11 false} {13 true}]
	ps = ps[2:3]
	fmt.Println(ps) // [{7 true}]
	ps = ps[0:]
	fmt.Println(ps) // [{7 true}]

	func printSlice(s []int) {
		fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	}

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // len=6 cap=6 [2 3 5 7 11 13] 

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s) // len=0 cap=6 []

	// Extend its length.
	s = s[:4]
	printSlice(s) // len=4 cap=6 [2 3 5 7]


	// Drop its first two values.
	s = s[2:]
	printSlice(s) // len=2 cap=4 [5 7]

	printSlice([] int{}) // [] == nil // len=0 cap=0 []

	// make creates a slice of a given length and capacity. The capacity is optional and defaults to the length.
	a := make([]int, 5)
	printSlice("a", a) // len=5 cap=5 [0 0 0 0 0]

	b := make([]int, 0, 5)
	printSlice("b", b) // len=0 cap=5 []

	c := b[:2]
	printSlice("c", c) // len=2 cap=5 [0 0]

	d := c[2:5]
	printSlice("d", d) // len=3 cap=3 [0 0 0]

	//append
	var stoke []int
	printSlice(stoke) // len=0 cap=0 []

	// append works on nil slices.
	stoke = append(stoke, 0)
	printSlice(stoke) // len=1 cap=1 [0]

	// The slice grows as needed.
	stoke = append(stoke, 1)
	printSlice(stoke) // len=2 cap=2 [0 1]

	// We can add more than one element at a time.
	stoke = append(stoke, 2, 3, 4)
	printSlice(stoke) // len=5 cap=6 [0 1 2 3 4]

	for i, v := range stoke {
		fmt.Printf("2**%d = %d\n", i, v) // 2**0 = 0, 2**1 = 1, 2**2 = 2, 2**3 = 3, 2**4 = 4
	}

}
