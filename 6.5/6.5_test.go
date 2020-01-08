package main

import ()

func ExampleInfo() {
	circ := circle{
		radius: 12.345,
	}

	squa := square{
		length: 15,
	}

	info(circ)
	info(squa)
	// Output:
	// 478.7756573542473
	// 225
}