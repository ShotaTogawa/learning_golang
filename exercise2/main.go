package main

import "fmt"

const (
	_ = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	s := "Hello world"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}
	fmt.Println("")


	x := 4
	fmt.Printf("%d\t%b\n", x, x)
	// bit shifting
	y := x << 1
	fmt.Printf("%d\t%b\n", y, y)



	fmt.Printf("%d\t%b\n", kb, kb)
	fmt.Printf("%d\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)

	fmt.Println("--- Exercises from here ----")
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
	exercise6()

}

func exercise1() {
	num := 100
	fmt.Printf("%d\t%b\t%#x\n", num, num, num)
}

func exercise2() {
	a := (42 == 42)
	b := (42 >= 10)
	c := (42 != 42)
	d := (42 < 10)
	e := (42 > 10)
	fmt.Println(a, b, c, d, e)
}

func exercise3() {
	const a = "apple"
	const b string = "orange"
	fmt.Println(a, b)
}

func exercise4() {
	a := 2056
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
	b := a << 1
	fmt.Printf("%d\t%b\t%#x\n", b, b, b)
}

func exercise5() {
	var a string = `hello world
	it's kinda sleepy because
	too early for me
`
	fmt.Println(a)
}

func exercise6() {
	const (
		_ = iota
		first = iota + 2017
		second = iota + 2017
		third = iota + 2017
		fourth = iota + 2017
	)

	fmt.Println(first, second, third, fourth)
}