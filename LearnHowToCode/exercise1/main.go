package main

import "fmt"

type myType int
var e  myType

func main() {
	// Exercise1 Short declaration
	x := 42
	y := "James Bond"
	z := true
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(x, y, z)

	// Exercise2/3 Var Declaration
	var a int = 42
	var b string = "Barry Bonds"
	var c bool = true

	fmt.Println(a, b, c)

	// Exercise3 Var Declaration

	s := fmt.Sprintf("%v\t%v\t%v", a, b, c)
	fmt.Println(s)

	// Exercise4 make own type
	var d myType;
	fmt.Println(d)
	fmt.Printf("%T\n", d)
	d = 42
	fmt.Println(d)

	// Exercise5
	e := int(d)
	fmt.Println(e)
	fmt.Printf("%T\n",e)






}
