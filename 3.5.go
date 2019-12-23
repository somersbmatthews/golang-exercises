/*
Hands-on exercise #5
Print out the remainder (modulus) which is found for each number between 10 and 100 when it is divided by 4.
*/

package main

import (
	"fmt"
)

func main() {
	rem := 0
	for i := 10, i <= 100, i++ {
		rem = i % 4
		fmt.Println(rem)
	}
}