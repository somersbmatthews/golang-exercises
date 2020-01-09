package main

import ()

func ExamplePrintAnonFunction() {
	printanonfunc()
	// Output:
	// 0
	// 1
	// 2
	// func()
	// 0
	// int
	// 0
	// 1
	// 2
	// this is g func()
	// done
}