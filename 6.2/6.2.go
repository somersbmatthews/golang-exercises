package main

import (
	"fmt"
)

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := Foo(ii...)
	fmt.Println(n)

	ii2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n2 := Bar(ii2)
	fmt.Println(n2)
}

func Foo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func Bar(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}



