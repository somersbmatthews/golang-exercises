package main

import (
	"fmt"
)

func main() {
	ListNums()
}

func ListNums() {
	func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
	}()

	fmt.Println("done")
}

