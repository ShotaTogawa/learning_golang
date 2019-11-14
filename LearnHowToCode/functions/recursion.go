package main

import "fmt"

//recursion is when a function calls itself

func main() {
	// 4 * 3 * 2 * 1
	fmt.Println(factorial(4))
	fmt.Println(factorial2(4))

}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	fmt.Println(n)
	return n * factorial(n - 1)
}

// return 4 * factorial(4 - 1)
// return 3 * factorial(3 - 1)
// return 2 * factorial(2 - 1)
// return 1 * factorial(1 - 1)
// return 1ddddd

func factorial2(num int) int {
	total := 1
	for ; num > 0; num-- {
		total *= num
	}
	return total
}