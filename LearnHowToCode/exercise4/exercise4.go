package main

import "fmt"

func main() {
	//array()
	//slice()
	//ex_map()
	//exercise1()
	//exercise2()
	//exercise3()
	//exercise4()
	//exercise5()
	//exercise6()
	//exercise7()
	exercise8_9()
}

func array() {
	// 基本GoではArrayではなく、sliceを使う
	// 型は一つの、決まった数の連続した要素
	// 配列の格納数を決定
	var x [5]int
	fmt.Println(x)      // [0,0,0,0,0]
	x[4] = 42           // index4に値を格納
	fmt.Println(x)      // [0,0,0,0,42]
	fmt.Println(len(x)) // 5
}

func slice() {
	// sliceは同じ型の値を格納できる。 a SLICE allows to
	// x := type{values}  composite literal
	// Composite literal construct values for structs, arrays, slices,
	// and maps and create a new value each time they are evaluated.
	// They consist of the type of the literal followed by a brace-bound list for elements
	x := []int{4, 5, 6, 7, 8, 42}
	y := []string{"apple", "orange", "banana"}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[0])
	fmt.Println(x[1:])
	fmt.Println(x[1:3])
	fmt.Println(x[:])

	for index, value := range x {
		fmt.Println(index, value)
	}

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}

	for _, v := range y {
		fmt.Println(v)
	}

	// append: func append(slice []T, elements ...T) [] T
	// ...はunlimited parameter
	// where T is a placeholder for any given type
	x = append(x, 10, 11, 12)
	fmt.Println(x)
	z := []int{234, 567, 890}
	x = append(x, z...)
	fmt.Println(x)

	// delete

	x = append(x[:2], x[6:]...)
	fmt.Println(x)

	// make([]type, length, capacity)

	c := make([]int, 10, 100)
	fmt.Println(c)
	fmt.Println(len(c))
	c = append(c, 1, 2, 4, 6)
	fmt.Println(cap(c))
	fmt.Println(c)

	// multi dimensional slice
	jb := []string{"James", "Bond", "Chocolate", "Banana"}
	mp := []string{"Miss", "MoneyPenny", "Strawberry", "Peanut"}

	xp := [][]string{jb, mp}
	fmt.Println(xp)
	fmt.Println(xp[0][1])
}

func ex_map() {
	// map[Type]Type
	// key, value data structure
	// unordered list
	// keyがない値は、default valueがはいる
	m := map[string]int{
		"James":           32,
		"Miss MoneyPenny": 27,
	}
	fmt.Println(m)
	v, ok := m["James"]
	fmt.Println(v, ok)

	// idiom
	if v, ok := m["James"]; ok {
		fmt.Println(v)
	}

	m["todd"] = 33

	for key, val := range m {
		fmt.Println(key, val)
	}

	xi := []int{4, 5, 6, 7, 8}

	for i, v := range xi {
		fmt.Println(i, v)
	}

	// delete
	// keyがない場合でも、エラーにならない
	delete(m, "todd")
	fmt.Println(m)

	if v, ok := m["James"]; ok {
		fmt.Println(v)
		delete(m, "James")
	}
	fmt.Println(m)
}

func exercise1() {
	a := [5]int{23, 56, 78, 96, 14}
	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", a)
}

func exercise2() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", a)
}

func exercise3() {
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(a[:5])
	fmt.Println(a[5:])
	fmt.Println(a[2:7])
	fmt.Println(a[1:6])
}

func exercise4() {
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	a = append(a, 52)
	fmt.Println(a)
	a = append(a, 53, 54, 55)
	fmt.Println(a)
	b := []int{56, 57, 58, 59}
	a = append(a, b...)
	fmt.Println(a)
}

func exercise5() {
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	a = append(a[:3], a[6:]...)
	fmt.Println(a)
}

func exercise6() {
	x := make([]string, 50, 50)
	x = []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`,
		` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`,
		` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`,
		` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`,
		` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
		` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`,
		` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`,
		` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`,
		` Wisconsin`, ` Wyoming`}

	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}

}

func exercise7() {
	x := []string{"James", "Bond", "Shaken"}
	y := []string{"Miss", "MoneyPenny", "Helloo"}
	z := [][]string{x, y}
	fmt.Println(z)
	for i, v := range z {
		fmt.Println("record: ", i, v)
		for j, v2 := range v {
			fmt.Printf("\t index posisiton: %v \t value: \t %v", j, v2)
		}
	}
}

func exercise8_9() {
	m := map[string][]string{
		`bond_james`: []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`: []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	fmt.Println(m)

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

	m["Card"] = []string{"Heart", "Dia", "Spade", "Clover"}
	fmt.Println(m)

	delete(m, "no_dr")
	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
