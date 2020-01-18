package main

import "fmt"

type person struct {
	first string
}

type human interface {
	Speak()
}

func (p *person) Speak() {
	fmt.Println("Hello")
}

func saySomething(h human) {
	h.Speak()
}

func main() {
	p1 := person{
		first: "James",
	}

	saySomething(&p1)

	// does not work
	// saySomething(p1)

	p1.Speak()
}
