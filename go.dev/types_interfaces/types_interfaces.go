package main

import "fmt"

// Structs provide a way to group related data together into a single object
// while interfaces enable defining behavior without enforcing any particular implementation

type Person struct {
	Name string
	Age  int
}

// Structs can also contain methods associated with them.
// These methods can be used to perform operations on the data encapsulated
// by the struct.
func (p *Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

// we have defined a Messenger interface with a single method signature SendMessage().
// Any type that implements this method can be considered as satisfying the Messenger interface.
type Messenger interface {
	// The SendMessage method takes a string message as input and returns a string.
	SendMessage(message string) string
}

// Let's consider an implementation of the Person struct
// that satisfies the Messenger interface:
func (p *Person) SendMessage(message string) string {
	fmt.Printf("%s says: %s\n", p.Name, message)
	return p.Name
}

func main() {
	// we define a method called Greet() associated with the Person struct which prints
	// a greeting message along with the person's name.
	person := Person{Name: "Alice", Age: 30}
	animal := Person{Name: "James", Age: 5}
	person.Greet()
	// we created an instance of our structure and assigned it to a variable of type Messenger.
	// We then called the method on this variable which was implemented in our struct.
	var messenger Messenger = &person
	var sender Messenger = &animal
	messenger.SendMessage("Hello, World!")
	sender.SendMessage("Woof, Woof!")
}
