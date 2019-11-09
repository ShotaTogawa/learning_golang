package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("my first func expression")
	}

	g := func(x int) {
		fmt.Println("g is", x)
	}
	f()
	g(30)
}
