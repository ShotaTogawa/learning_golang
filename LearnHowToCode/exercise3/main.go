package main

import (
	"fmt"
)

func main() {
	// for init; condition; post {}
	for i := 0; i <= 10; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("The outer loop: %d\t The inner loop: %d\t\n",i, j)
		}
	}

	x := 1
	// for range {}
	for { if x > 9 {
			break
		}
		fmt.Println(x)
		x++
	}
	fmt.Println("done")

	for a := 33; a <= 122; a++ {
		fmt.Printf("%v\t%x\t%#U\n", a, a, a)
	}

	// 二つのstatementを１列に書くときに;がいる
	if x == 1 {
		fmt.Printf("it's 1")
	} else if x == 0 {
		fmt.Println("it's 0")
	} else {
		fmt.Printf("it's not 1"); fmt.Println("try again")
	}

	// switch
	switch "Bond"{
	case "Money":
		fmt.Println("this should nt print")
	case "Q":
		fmt.Println("this should nt print2")
	case "Bond":
		fmt.Println("print")
		//fallthrough // returnがtrueになっても、次のケースを実行する
	default:
		fmt.Println("default")
	}

	exercise1()
	exercise2()
	//exercise3()
	//exercise4()
	//exercise5()
	//exercise6()
	//exercise7()

}

func exercise1 () {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}

func exercise2 () {
	for i := 65; i < 90; i++ {
		fmt.Println(i)
		for j :=0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}

}

func exercise3 () {
	year := 1992
	for year <= 2019 {
		fmt.Println(year)
		year++
	}
}

func exercise4 () {
	bd := 1992
	for {
		if bd > 2019 {
			break
		}
		fmt.Println(bd)
		bd++
	}
}

func exercise5 () {
	for i:= 10; i <= 100; i++ {
		fmt.Printf("When %v is divided by 4, the remander (aka module) is %v\n",i, i%4)
	}
}

func exercise6 () {
	x := 1
	if x < 10 {
		fmt.Println("x is under 10")
	} else if x == 10 {
		fmt.Println("x is 10")
	} else {
		fmt.Println("x is over 10")
	}
 }

func exercise7 () {

	x := 1
	switch  {
	case x < 10:
		fmt.Println("x is under 10")
	case x == 10:
		fmt.Println("x is 10")
	case x > 10:
		fmt.Println("x is 10")
	default:
		fmt.Println("x is not a number")
	}
}