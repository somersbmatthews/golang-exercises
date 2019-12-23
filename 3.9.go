/*
Hands-on exercise #9
Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.
solution:
*/


package main

import (
	"fmt"
)

var favSport string = "golf"

func main() {
	switch {
	case "tennis":
		fmt.Println("should not print")
	case "golf": 
		fmt.Println("golf is my favorite sport")
	}
}

