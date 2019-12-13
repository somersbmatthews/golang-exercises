package main

import "fmt"

type megatype int

var x megatype
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x, "is value for x")
	y = int(x)
	fmt.Printf("%T", y)
	fmt.Println(" is type of y")
	fmt.Println(y, "is value for y")
}
