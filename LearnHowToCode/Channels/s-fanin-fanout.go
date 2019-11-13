package main

import "fmt"

func producer1(first chan int) {
	defer close(first)
	for i := 0; i < 10; i++ {
		// 一周ごとに、iの値がfirstに送信される
		first <- i
	}
}

// 以下の引数のように、明示的にチャネルをかける。
func multi(first <-chan int, second chan<- int) {
	defer close(second)
	// ここでfirstの値が受信される。
	for i := range first {
		// 受信したfirstの値に2をかけて、secondに送信する
		second <- i * 2
	}
}

func multi4(second <-chan int, third chan<- int) {
	defer close(third)
	// secondに送信された値を受信する。
	for i := range second {
		// 受信したsecondの値に4をかけて、thirdに送信する
		third <- i * 4
	}
}

func main() {
	first := make(chan int)
	second := make(chan int)
	third := make(chan int)

	go producer1(first)
	go multi(first, second)
	go multi4(second, third)
	for result := range  third{
		// この場合、10周した値がここに入る
		fmt.Println(result) //0 8 16 24 32 40 48 56 64 72
	}
}