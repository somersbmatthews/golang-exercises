package main

import ()

func ExampleSpeak() {
	p1 := Person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p1.Speak()
	// Output: I am James Bond, and I am 32 years old.

}
