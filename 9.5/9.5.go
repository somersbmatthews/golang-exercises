package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	AtomicRun()
}

func AtomicRun() {
	var wg sync.WaitGroup
	var incrementer int64

	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			newValue := atomic.AddInt64(&incrementer, 1)
			println(newValue)

			wg.Done()
			runtime.Gosched()
		}()
	}
	wg.Wait()
	fmt.Println("end value", incrementer)

}
