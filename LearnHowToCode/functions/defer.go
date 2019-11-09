package main

import "fmt"

func main() {
	defer goo()
	gle()
}

func goo() {
	defer func(){
		fmt.Println("defer ran")
	}()
	fmt.Println("goo")
}

func gle() {
	fmt.Println("gle")
}