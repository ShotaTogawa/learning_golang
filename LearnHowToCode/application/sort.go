package main

import (
	"fmt"
	"sort"
)

type animal struct {
	first string
	age int
}

type ByAge []animal
type ByName []animal

// Custom Sort
func (a ByAge) Len() int {return len(a)}
func (a ByAge) Swap(i, j int) {a[i], a[j] = a[j], a[i]}
func (a ByAge) Less(i, j int) bool {return a[i].age < a[j].age}

func (bn ByName) Len() int {return len(bn)}
func (bn ByName) Swap(i, j int) {bn[i], bn[j] = bn[j], bn[i]}
func (bn ByName) Less(i ,j int) bool {return bn[i].first < bn[j].first}

func main() {
	xi := []int{4,7,20,1,5,70,13,5}
	si := []string{"banana", "apple", "orange", "grape"}

	// Sort
	// 再代入する必要はない。sortをかけた時点で、slice内のデータがソートされる。
	sort.Ints(xi)
	sort.Strings(si)
	fmt.Println(xi) // [1 4 5 5 7 13 20 70]
	fmt.Println(si) // [apple banana grape orange]

	a1 := animal{"dog", 6}
	a2 := animal{"cat", 4}
	a3 := animal{"rabbit", 8}
	a4 := animal{"rat", 3}

	animals := []animal{a1,a2,a3,a4}
	fmt.Println(animals)
	sort.Sort(ByAge(animals))
	fmt.Println(animals)

	sort.Sort(ByName(animals))
	fmt.Println(animals)

}
