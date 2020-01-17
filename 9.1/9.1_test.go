package main

import ()

func ExampleMakeGoRoutine() {
	MakeGoRoutine()
	// Output:
	// begin cpu 8
	// begin gs 1
	// mid cpu 8
	// mid gs 3
	// hello from thing two
	// hello from thing one
	// about to exit
	// end cpu 8
	// end gs 1
}
