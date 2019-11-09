package main

import "fmt"

// callback
// - passing a func as an argument

func main() {
	fmt.Println(sum(1,2,3,4,5,6,7,8))
	ii := []int{5,6,7,8,9,10,11,12,13,14}
	s := sum(ii...)
	fmt.Println(s)
	fmt.Println(even(sum, ii...))
	fmt.Println(odd(sum, ii...))
}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v % 2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func odd(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, x := range vi {
		if x % 2 == 1 {
			yi = append(yi, x)
		}
	}
	return  f(yi...)
}