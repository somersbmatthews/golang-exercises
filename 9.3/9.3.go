package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	incrementRoutine()
}

func incrementRoutine() {
	var wg sync.WaitGroup

	incrementer := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			fmt.Println(incrementer)
			wg.Done()
		}()
	}

	// var num int = 1
	// var sum int = 0

	// func increment() {
	// 	sum = sum + num
	// }

	// func runRoutine() {
	// 	increment(num)
	// 	fmt.Println(num)
	// 	fmt.Println(sum)
	// }
}
