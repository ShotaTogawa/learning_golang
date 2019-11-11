package main

import (
	"encoding/json"
	"fmt"
)

// func Marshal(v interface{}) ([]byte, error)
	// Marshal returns the JSON encoding of v.
	// I just make the fields upper case with Marshall. It will grab the fields and export them

// func Unmarshal(data []byte, v interface{}) error
	// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v

type person struct {
	First string
	Last string
	Age int
}

type human struct {
	First string `json:"First"`
	Last string `json:"Last"`
	Age int `json:"Age"`
}


func main() {
	p1 := person{
		First: "James",
		Last: "Bond",
		Age: 20,
	}

	p2 := person{
		First: "Miss",
		Last: "Money Penny",
		Age: 29,
	}

	people := []person{p1, p2,}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil{
		fmt.Println(err)
	}
	// byteをstringに戻す
	fmt.Println("json",string(bs))


	s := `[{"First":"James","Last":"Bond","Age":20},{"First":"Miss","Last":"Money Penny","Age":29}]`
	bs2 := []byte(s)
	fmt.Printf("%T\n",s) // string
	fmt.Printf("%T\n", bs2) // []uint8

	var humans []human

	// 第一引数はbinary, 第二引数はアドレスを渡す
	err = json.Unmarshal(bs2, &humans)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("all of the data", humans)

	for i, v := range humans {
		fmt.Println("\nHuman number", i)
		fmt.Println(v.First, v.Last, v.Age)
	}


}


