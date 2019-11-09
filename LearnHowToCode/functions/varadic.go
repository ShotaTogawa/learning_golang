package main

import "fmt"

func main() {
	x := foo(2,3,4,5,6,7,8,9)
	fmt.Println(x)
}

// func (r receiver) identifier (params(s)) (return(s)) { ... }
// varadic parameter "..."
// slice of int. ...を引数にした際、スライスになる
func foo (x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position,", i, "we are now adding",v ,"to the total which is now")
	}
	fmt.Println("The total is, ", sum)
	return sum
}
