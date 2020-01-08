package main_test

import (
	"golang-exercises/6.3"
	"fmt"
)

func ExampleFoo() {
	defer main.Foo()
	fmt.Println("Hello, playground")
	// Output:
	// Hello, playground
	// foo ran
	// foo DEFER ran
}





