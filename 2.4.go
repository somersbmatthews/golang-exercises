package main

import "fmt"

const num = 42

var shifted_num int

func main() {
	fmt.Printf("%d\t%b\t%#x\n", num, num, num)
	shifted_num = num << 1
	fmt.Printf("%d\t%b\t%#x\n", shifted_num, shifted_num, shifted_num)
}
