/*
One of the Go's differentiating language is how it treat interfaces.
The usual usual requirements that TYPE specify what interface to implement is not present in GO,
making it one of its most powerful concepts to grasp early in your learning journey.

We start by defining an interface TYPE to work with we declare humanoid with interface
*/

package main

import (
	"fmt"
)

// define an humanoid interface with MUST have methods: speak and walk methods returning strings
type humanoid interface {
	speak()
	walk()
}

// define a person TYPE that implements humanoid interface
// PS: we dont need to specify that person implements humanoid anywhere in the code.
type person struct{ name string }

// simply implementing the interface's methods will do the trick.
func (p person) speak() { fmt.Printf("%s is speaking...\n", p.name) }

func (p person) walk() { fmt.Printf("%s is walking...\n", p.name) }

// let's change it by implement Stringer interface at person struct
// person output is changed by simply implementing the Stringer interface
func (p person) String() string {
	return fmt.Sprintf("Hello! My name is %s", p.name)
}

// let's define a dog but only implement the walk method on it
type dog struct{}

func (d dog) walk() { fmt.Println("Dog is walking...\n") }

// implement the stringer interface for the person type

// define a dog TYPE that can walk but not speak

func main() {
	// invoke with a person
	p := person{name: "Lee"}
	// as expected, person value satisfies humanoid interface requirement.
	doHumanThings(p)

	// can we invoke with a dog?
	// d := dog{}
	// No, dog value failed to meet hte requirement even though it is partially met them
	// The key takeaway here is that TYPE are required to implement ALL of the methods of an interface.
	// TYPEs are allowed to implement as many interfaces as they wish.
	//doHumanThings(d)

	// This being the default output for person {Lee}
	fmt.Println(p)
}

func doHumanThings(h humanoid) {
	h.speak()
	h.walk()
}
