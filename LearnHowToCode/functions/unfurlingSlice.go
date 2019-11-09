package main

import "fmt"

func main() {
	xi := []int{1,2,3,4,5,6,7,8}
	// sliceを引数に渡す際には、"変数名..."
	x := unfurling("", xi...)
	fmt.Println(x)

	greeting("Hello", "Joe", "Anna", "Eileen")
}

// func (r receiver) identifier (params(s)) (return(s)) { ... }
// unfurling a slice (not a official word)

func unfurling (s string,x ...int) int {

	fmt.Println(x)
	fmt.Printf("%T\n", x) // []int
	fmt.Println(cap(x)) // 0
	fmt.Println(len(x)) // 0

	s1 := s
	fmt.Println(s1)
	sum := 0
	return sum
}

func greeting(prefix string, who ...string)   {
	if len(who) == 0 {
		fmt.Println("nobody")
	}
	for _, x := range who {
			fmt.Println(prefix, x)
		}
}

