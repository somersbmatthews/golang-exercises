package main

import (
	"fmt"
)

func main() {
	defer Foo()
	fmt.Println("Hello, playground")
}

func Foo() {
	defer func() {
		fmt.Println("foo DEFER ran")
	}()
	fmt.Println("foo ran")
}