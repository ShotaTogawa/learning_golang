package main

import (
	"fmt"
)

// In Go, values can be of more than one type.
// An interface allows a value to be of more than one type.
// We create an interface using this syntax: “keyword identifier type”
// so for an interface it would be: “type human interface”
// We then define which method(s) a type must have to implement that interface.
// If a TYPE has the required methods, which could be none (the empty interface denoted by interface{}),
// then that TYPE implicitly implements the interface and
// is also of that interface type.
// In Go, values can be of more than one type.

type person1 struct {
	first string
	last  string
}

type secretAgent struct {
	person1
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")
}

func (p person1) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

type human interface {
	speak()
}

func bar1(h human) {
	switch h.(type) {
	case person1:
		fmt.Println("I was passed into barrrrrr", h.(person1).first)
	case secretAgent:
		fmt.Println("I was passed into barrrrrr", h.(secretAgent).first)
	}
	fmt.Println("I was passed into bar", h)
}

type hotdog int

func main() {
	sa1 := secretAgent{
		person1: person1{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person1: person1{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	p1 := person1{
		first: "Dr.",
		last:  "Yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	bar1(sa1)
	bar1(sa2)
	bar1(p1)

	// conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
