package main_test

import (
	"golang-exercises/6.2"
	"fmt"
)

func ExampleFoo() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(main.Foo(ii...))
	// Output: 45
}

func ExampleBar() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(main.Bar(ii))
	// Output: 45
}