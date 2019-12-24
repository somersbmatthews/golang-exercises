/*
Hands-on exercise #3
Using the code from the previous example, use SLICING to create the following new slices which are then printed:
*/

package main

import (
	"fmt"
)

func main() {
	s1 := []int{42, 43, 44, 45, 46}
	s2 := []int{47, 48, 49, 50, 51}
	s3 := []int{44, 45, 46, 47, 48}
	s4 := []int{43, 44, 45, 46, 47}

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
}

