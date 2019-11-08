package main

import "fmt"

// A struct is a data structure which allows us to compose together values of different types
// Within a struct, non-blank field names must be unique

// we create VALUES of a certain TYPE that are stored in VARIABLES
// and those VARIABLES have IDENTIFIED
// VARIABLES IDENTIFIED TYPE

// clear and readable

type person struct {
	first string
	last string
	age int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {

	//exercise1()
	exercise3()
	exercise4()
}


func leacture() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last: "Bond",
			age: 28,
		},
		ltk: true,
	}

	p2 := person{
		first: "Miss",
		last: "MoneyPenny",
	}

	fmt.Println(sa1)
	fmt.Println(sa1, p2)
	// sa1.person.firstとsa1.firstは同じものをさす。
	// firstが複数structの中に存在する場合に使い分ける
	fmt.Println(sa1.person.first,sa1.age, sa1.last)
	fmt.Println(p2.first,p2.age, p2.last)


	// Anonymous struct
	p3 := struct {
		first string
		last string
		age int
	}{
		first: "Tom",
		last: "Brown",
		age: 30,
	}
	fmt.Println(p3)

	// constant
	const bar int = 100
	fmt.Println(bar)

}


type human struct{
	first_name string
	last_name string
	ice_cream []string
}

func exercise1() {

	p1 := human{
		first_name: "tom",
		last_name: "brown",
		ice_cream: []string{"chocolate","Vanilla"},
	}

	p2 := human{
		first_name: "Jams",
		last_name: "Bond",
		ice_cream: []string{"banana","strawberry"},
	}

	m := map[string]human {
		p1.last_name: p1,
		p2.last_name: p2,
	}

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v.first_name)
	}

	fmt.Println(m)

	fmt.Println(p1.first_name, p1.last_name)
	for _ , v := range p1.ice_cream {
		fmt.Println(v)
	}
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func exercise3() {

	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel:true,
	}
	fmt.Println(t.doors)

	s := sedan {
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: true,
	}

	fmt.Println(s.doors)

}

func exercise4() {
	someone := struct {
		book string
		favorite map[int]string
		favDrink []string
	}{
		book: "Harry potter",
		favorite: map[int]string{
			10: "game",
			20: "english",
			30: "programming",
		},
		favDrink: []string{
			"coffee", "latte", "water",
		},
	}
	fmt.Println(someone.book)
	fmt.Println(someone.favorite[10])
	fmt.Println(someone.favDrink[1])

}