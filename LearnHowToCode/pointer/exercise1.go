package main

import "fmt"

type men struct {
	name string
}

func changeMe(m *men) {
	m.name = "Miss MoneyPenny"
	// (*m).name = "Miss MoneyPenny"
}

func main() {
	c := 10
	fmt.Println(&c)
	m := men{
		name: "James Bond",
	}

	fmt.Println(m)
	changeMe(&m)
	fmt.Println(m)

}