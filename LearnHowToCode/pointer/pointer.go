package main

import "fmt"

// pointer is a just like pointing to some location in memory where a value is stored
// *はポインターをさす。
// * is an operator. this * right here is part of type.
// *int はポインター

// when to use
//

func main() {
	a := 42
	fmt.Println(a)
	// メモリのアドレスを表示
	fmt.Println(&a) // &はアドレスをさす
	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n", &a)
	b := &a
	fmt.Println(b) // aと同じアドレスをさす
	fmt.Println(*b) // * gives you the value stored at an address when you have the address
	fmt.Println(*&a) // 42

	c := "apple"
	fmt.Println(&c)
	fmt.Println(*&c)
	d := &c
	fmt.Println(*d)

	// 131. when you use pointer
	x := 0
	fmt.Println("x before", x)
	fmt.Println("x before", &x)
	foo(&x)
	fmt.Println("x after", x)
	fmt.Println("x after", &x)

}

func foo(y *int) {
	fmt.Println("y before", y)
	fmt.Println("y before", *y)
	*y = 43
	fmt.Println("y after", y)
	fmt.Println("y after" ,*y)
}