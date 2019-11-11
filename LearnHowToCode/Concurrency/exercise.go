package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

type person struct {
	first string
}

type human interface {
	speak()
}


func main() {
	exercise1()

	// exercise2
	p1 := person{first:"Tom"}
	saySomething(&p1)
	p1.speak()

	//exercise3()
	exercise5()

	exercise6()
}


func (p *person) speak() {
	fmt.Println("Hello")
}

func saySomething(h human) {
	h.speak()
}


func exercise1() {

	fmt.Println("s-CPUs:",runtime.NumCPU())
	fmt.Println("s-goroutines:", runtime.NumGoroutine())
	var wg1 sync.WaitGroup

	wg1.Add(2)

	go func() {
		fmt.Println("hello from thing one")
		wg1.Done()
	}()
	go func() {
		fmt.Println("hello from thing two")
		wg1.Done()
	}()

	fmt.Println("before wg1.Wait()")

	fmt.Println("m-CPUs:",runtime.NumCPU())
	fmt.Println("m-goroutines:", runtime.NumGoroutine())

	wg1.Wait()

	fmt.Println("after wg1.Wait()")
	fmt.Println("e-CPUs:",runtime.NumCPU())
	fmt.Println("e-goroutines:", runtime.NumGoroutine())


}

func exercise3() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	incrementer := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func(){
			mu.Lock()

			v := incrementer
			//runtime.Gosched()
			v++
			incrementer = v
			fmt.Println("middle", incrementer)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end",incrementer)

}

func exercise5() {
	var wg sync.WaitGroup
	var incrementer int64

	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func(){
			atomic.AddInt64(&incrementer, 1)
			fmt.Println(atomic.LoadInt64(&incrementer))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end",incrementer)

}

func exercise6() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}