package main

import (
	"fmt"
)

type Person struct {
	first string
	last  string
	age   int
}

func (p Person) Speak() {
	fmt.Printf("I am %s %s, and I am %d years old.", p.first, p.last, p.age)
}

func main() {
	p1 := Person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p1.Speak()
}