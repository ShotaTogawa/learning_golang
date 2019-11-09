package main

import "fmt"

func main() {
	bar("James")
	s1 := woo("Bond")
	fmt.Println(s1)
	a, b := person("tanaka", "taro")
	fmt.Println(a)
	fmt.Println(b)
}

// func (r receiver) identifier(parameters) (return(s)){...}
// Everything in Go is Pass by Value
func bar(s string) {
	fmt.Println("Hello ",s)
}

func woo(s1 string) string {
	return fmt.Sprint("Hello from woo, ", s1)
}

func person(fs string, ls string )(string, bool) {
	a := fmt.Sprint(fs, ls, `, say Hello`)
	b := true
	return a, b
}