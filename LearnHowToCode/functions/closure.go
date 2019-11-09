package main

import "fmt"

// closure is a concept of we want to close.
// the scope of a variable and contain  it to certain areas who are going to enclose some code
// scope of the valuable is limited

func main() {
	// a and b have different scopes they have different memory locations where that values is stored
	a := incrementor()
	b := incrementor()
	fmt.Println(a()) // 1
	fmt.Println(a()) // 2
	fmt.Println(a()) // 3
	fmt.Println(a()) // 4
	fmt.Println(b()) // 1
	fmt.Println(b()) // 2
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
