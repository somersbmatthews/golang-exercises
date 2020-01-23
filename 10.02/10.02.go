package main

import (
	"fmt"
)

func main() {
	MakeChannel()
}

func MakeChannel() {
	c := make(chan int)
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)
}
