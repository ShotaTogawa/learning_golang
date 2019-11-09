package main

import "fmt"

type myType int

var e myType

func main() {
	exercise1()
	exercise2And3()
	exercise4And5()

}

type bucks int

var z int = 0

func lecture() {
	// 最初は:がいる
	// short declaration operator
	// Declare a variable and assign a value
	// function内でしか下記の定義はできない
	x := 42
	fmt.Println(x)
	// 再定義は:がいらない
	x = 99
	fmt.Println(x)

	// 下記はグローバルに定義できる
	var y = 10

	fmt.Println(y)
	fmt.Printf("%T\n", y)
	// %dはDecimal, %bはBinary

	// Goは自分でタイプを作れる。
	bucks := 10
	fmt.Println(bucks)

	// %dはDecimal, %bはBinary
}

func exercise1() {
	// Exercise1 Short declaration
	x := 42
	y := "James Bond"
	z := true
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(x, y, z)

}

func exercise2And3() {
	// Exercise2/3 Var Declaration
	var a int = 42
	var b string = "Barry Bonds"
	var c bool = true

	fmt.Println(a, b, c)

	// Exercise3 Var Declaration

	s := fmt.Sprintf("%v\t%v\t%v", a, b, c)
	fmt.Println(s)

}

func exercise4And5() {
	// Exercise4 make own type
	var d myType
	fmt.Println(d)
	fmt.Printf("%T\n", d)
	d = 42
	fmt.Println(d)

	// Exercise5
	e := int(d)
	fmt.Println(e)
	fmt.Printf("%T\n", e)

}
