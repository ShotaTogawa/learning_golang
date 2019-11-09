package main

import "fmt"

type person struct {
	first string
	last string
}

type company struct {
	person
	it bool
}

type teacher struct {
	first string
	last string
	country string
	gender string
}

type school struct {
	teacher
	students int
}

func (t teacher) teach() string {
	return "teaching"
}


// Method クラスのメソッドのように使える
// 自分で定義したタイプに対して、メソッドを定義できる。
func (c company) speak() {
	fmt.Println("I am", c.first, c.last)
}

func main() {
		comp := company{
			person: person{
				first: "Bob",
				last: "Brown",
			},
			it:     true,
		}

		fmt.Println(comp)
		comp.speak()

		s1 := school{
			teacher:  teacher{
				first: "tanaka",
				last: "taro",
				country: "Japan",
			},
			students: 20,
		}

		fmt.Println(s1.first, s1.teach(), "to", s1.students )

}
