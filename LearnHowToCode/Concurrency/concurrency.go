package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Go programming language is the first major software programming language ever to
// have been written ever build to natively take advantage of multiple cores

// concurrency vs parallelism

// In programming, concurrency is the composition of independently executing processes,
// while parallelism is the simultaneous execution of (possibly related) computations.
// Concurrency is about dealing with lots of things at once. Parallelism is about doing lots of things at once.

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t",  runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	bar()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()

}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}