package main

import "fmt"

func main() {
	fmt.Println(hello())
	x := bar2()
	fmt.Printf("%T\n", x) // func int

	// 最初の()がfuncの呼び出しで、二つ目が実行
	y := bar2()()
	fmt.Println(y)
}

func hello() string {
	s := "Hello"
	return s
}

// func identifier return { ... }

func bar2() func() int {
	return func() int {
		return 451
	}
}