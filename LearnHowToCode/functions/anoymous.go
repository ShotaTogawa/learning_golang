package main

import "fmt"

func main() {
	func() {
		fmt.Println("Anonymous func ran")
	}()

	func (x int) {
		fmt.Println("age is", x)
	}(42)
}
