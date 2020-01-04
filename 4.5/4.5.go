/*
Hands-on exercise #5
To DELETE from a slice, we use APPEND along with SLICING. For this hands-on exercise, follow these steps:
start with this slice
x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
use APPEND & SLICING to get these values here which you should ASSIGN to a variable “y” and then print:
[42, 43, 44, 48, 49, 50, 51]
*/

package main

import (
	"fmt"
)

func DeleteByAppend(s []int) []int {
	x := append(s[:3], s[6:]...)

	fmt.Println(x)
	return x
}
