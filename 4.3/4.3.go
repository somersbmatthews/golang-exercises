/*
Hands-on exercise #3
Using the code from the previous example, use SLICING to create the following new slices which are then printed:
[42 43 44 45 46]
[47 48 49 50 51]
[44 45 46 47 48]
[43 44 45 46 47]
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(Slicing())
	
}

func Slicing() ([]int, []int, []int, []int, []int) {
	
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	return x[:5], x[5:], x[2:7], x[1:6], x
}

