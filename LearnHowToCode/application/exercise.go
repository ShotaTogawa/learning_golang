package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type user struct {
	First string `json:First`
	Age   int `json:Age`
}

func main() {
	exercise1And2()
	exercise4()
	exercise5()
}

func exercise1And2() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	pj, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(pj))

	// exercise2
	s := `[{"First":"James","Last":"Bond","Age":20},{"First":"Miss","Last":"Money Penny","Age":29}]`
	var users1 []user
	bs2 := []byte(s)
	err = json.Unmarshal(bs2, &users1)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range users1 {
		fmt.Println(v.First, v.Age)
	}


	// exercise3
	err = json.NewEncoder(os.Stdout).Encode(users1)
	if err != nil {
		fmt.Println(err)
	}
}



func exercise4() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	// sort xi
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	// sort xs
	sort.Strings(xs)
	fmt.Println(xs)

}

type user2 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type SortAge []user2
type ByLastName []user2
type ByQuote []user2


func (a SortAge) Len() int {return len(a)}
func (a SortAge) Swap(i, j int) {a[i], a[j] = a[j], a[i]}
func (a SortAge) Less(i, j int) bool {return a[i].Age < a[j].Age}

func (bn ByLastName) Len() int {return len(bn)}
func (bn ByLastName) Swap(i, j int) {bn[i], bn[j] = bn[j], bn[i]}
func (bn ByLastName) Less(i ,j int) bool {return bn[i].Last < bn[j].Last}

//func (bn ByQuote) Len() int {return len(bn)}
//func (bn ByQuote) Swap(i, j int) {bn[i], bn[j] = bn[j], bn[i]}
//func (bn ByQuote) Less(i ,j int) bool {return bn[i].Sayings[i] < bn[j].Sayings[j]}
//

func exercise5() {
	u1 := user2{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user2{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user2{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user2{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	sort.Sort(SortAge(users))
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age )
		for _, q := range u.Sayings{
			fmt.Println("\t", q)
		}
	}

	fmt.Println("--------------")

	sort.Sort(ByLastName(users))
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age )
		for _, q := range u.Sayings{
			fmt.Println("\t", q)
		}
	}

	fmt.Println("--------------")

	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age )
		sort.Strings(u.Sayings)
		for _, q := range u.Sayings{
			fmt.Println("\t", q)
		}
	}
}

