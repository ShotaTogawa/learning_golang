package main

import "fmt"

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func goroutine2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{1,2,3,4,5,6}
	c := make(chan int)
	go goroutine1(s, c)
	go goroutine2(s, c)
	x := <-c
	fmt.Println(x)
	y := <-c
	fmt.Println(y)

	fmt.Println("--------")
	bufferdChannel()
	fmt.Println("--------")
	howToLoopChannel()

}

func bufferdChannel() {
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))
	// これ以上受信するとエラーになる、下記はエラー
	//ch <- 300 受信できない

	x := <-ch
	fmt.Println(x)

	// 上記のように一度取り出すと、再度受信できる
	ch <- 300
	fmt.Println(len(ch))

}

func howToLoopChannel() {
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))
	close(ch)

	for c := range ch{
		fmt.Println(c)
	}
}