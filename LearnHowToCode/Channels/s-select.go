package main

import (
	"fmt"
	"time"
)

// 同時にチャネルを受信できる
// selectはforがないと使えない

func routine1(ch chan<- string) {
	for {
		ch <- "packet from 1"
		time.Sleep(3 * time.Second)
	}
}

func routine2(ch chan<- int) {
	for {
		ch <- 100
		time.Sleep(1 * time.Second)
	}
}

func main() {

	select_default()

	fmt.Println("###########")

	c1 := make(chan string)
	c2 := make(chan int)

	go routine1(c1)
	go routine2(c2)

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

func select_default() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500* time.Millisecond)
	OuterLoop:
		for {
			select {
			case <- tick:
				fmt.Println("tick.")
				case <- boom:
					fmt.Println("Boom!")
					break OuterLoop
			default:
				fmt.Println("	.")
				time.Sleep(50 * time.Millisecond)
			}
		}
	fmt.Println("#########")
}
