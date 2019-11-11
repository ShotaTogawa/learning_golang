package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GOROUTINES:", runtime.NumGoroutine())

	//counter := 0

	// atomic
	var counter int64

	var wg sync.WaitGroup
	const gs = 100
	wg.Add(gs)

	//var mu sync.Mutex // race conditionを避けるために使う

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)

			//mu.Lock()

			//v := counter
			//time.Sleep(time.Second)

			runtime.Gosched()

			fmt.Println("Counter\t", atomic.LoadInt64(&counter))

			//v++
			//counter = v

			//mu.Unlock()
			wg.Done()
		}()
		fmt.Println("GOROUTINES:", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("GOROUTINES:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
