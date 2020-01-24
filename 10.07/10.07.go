package main

import (
	"fmt"
)

var channel = make(chan int)

func main() {
	LaunchRoutines()
	PrintNums()
	close(channel)
}

func LaunchRoutines() {
	for i := 1; i <= 10; i++ {
		go Routine()
	}
}

func Routine() {
	for i := 1; i <= 10; i++ {
		channel <- i
	}
}

func PrintNums() {
	for i := 1; i <= 100; i++ {
		fmt.Println(<-channel)
	}
}
