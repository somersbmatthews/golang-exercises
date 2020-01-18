package main

import ()

func ExampleSaySomething() {
	p1 := person{
		first: "James",
	}
	saySomething(&p1)
	// Output:
	// Hello
}
