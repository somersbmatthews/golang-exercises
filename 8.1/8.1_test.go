package main

import ()

func ExampleJsonMarshal() {
	JsonMarshal()
	// Output:
	// [{James 32} {Moneypenny 28} {M 54}]
	// [{"First":"James","Age":32},{"First":"Moneypenny","Age":28},{"First":"M","Age":54}]
}
