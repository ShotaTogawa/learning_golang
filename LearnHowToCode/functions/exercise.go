package main

import "fmt"

type man struct {
	first string
	last string
	age int
}

// method
func (m man) speak() {
	fmt.Println("I am", m.first, m.last, "and I am", m.age,"age years old")
}

func main() {
	i := exercise1_1()
	ii, s := exercise1_2()
	fmt.Println(i, ii, s)

	nums := []int{1,2,3,4,5,6,7,8,9}
	e2_1 := exercise2_1(nums...)
	e2_2 := exercise2_2(nums)
	fmt.Println(e2_1, e2_2)

	// exercise4
	ex4 := man{
		first: "tanaka",
		last:  "taro",
		age:   20,
	}
	ex4.speak()

	// exercise 6
	func() {
		fmt.Println("anonymous function")
	}()

	// exercise 7
	exercise7(ex7)

	// exercise 8
	ex8 := exercise8()
	ex8()

	// exercise 9

	g := func(xi []int) int {
			if len(xi)  == 0 {
				return 0
			}
			if len(xi) == 1 {
				return xi[0]
			}
		return xi[0] + xi[len(xi) - 1]
	}

	re := exercise9(g, []int{1,2,3,4,5,6,7,8})
	fmt.Println(re)

	// exercise 10
	ex10 := exercise10()
	fmt.Println(ex10())
	fmt.Println(ex10())
	fmt.Println(ex10())
	fmt.Println(ex10())

}

func exercise1_1() int {
	return 1
}

func exercise1_2() (int, string) {
	return 1, "hi"
}

func exercise2_1 (nums ...int) int {
	total := 0
	for _, n := range nums {
			total += n
	}
	return total

}

func exercise2_2 (num []int) int {
	total := 0
	for _, n := range num {
		total += n
	}
	return total
}

// exercise3はdefer.go内で実装
// exercise5は別ファイル

func ex7() {
	fmt.Println("foo")
}

func exercise7(f func()) {
	fmt.Println("callback")
	f()
}

func exercise8() func() {
	x := func() {
		fmt.Println("return function")
	}
	return x
}

func exercise9(f func(xi []int) int, ii []int) int {
	n := f(ii)
	n++
	return n
}

func exercise10() func() int {
	var init int
	return func() int {
		init++
		return init
	}
}

