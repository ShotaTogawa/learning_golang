package main

import "fmt"

//channels block

func main() {
	c := make(chan int, 2)
	cr := make(<-chan int) //receive
	cs := make(chan <- int) //send

	go func () {
		c <- 42
		c <- 43
	}()
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("-------")
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n",cr)
	fmt.Printf("%T\n", cs)

	// general to specific convert
	fmt.Println("-----")
	fmt.Printf("c\t%T\n", (<-chan int)(c))
	fmt.Printf("c\t%T\n", (chan <-int)(c))

	using_with_func()
}

func using_with_func() {
	c := make(chan int)
	// send
	go send(c)

	// receive
	receive(c)
}


// send
func send(c chan <- int) {
	c <- 42
}

// receive
func receive(c <-chan int) {
	fmt.Println("the value received from the channel:",<-c)
}