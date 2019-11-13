package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	// send
	go send1(even, odd, quit)

	// receive
	receive1(even, odd, quit)

	fmt.Println("about to exit")

}

func send1(e,o chan<- int,q chan <- bool) {
	for i := 0; i < 100; i++{
		if i % 2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
}

func receive1(e, o <-chan int, q <- chan bool) {
	for {
		select{
			case v := <-e:
				fmt.Println("from th eve channel", v)
			case v := <-o:
				fmt.Println("from the odd channel", v)
			case i, ok := <-q:
				if !ok {
					fmt.Println("from comma ok", i, ok)
					return
				} else {
					fmt.Println("from comma ok", i)
				}
		}
	}
}