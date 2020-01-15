package main

import ()

func ExampleJsonUnmarshal() {
	JsonUnmarshal()
	// Output:
	// [{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]
	// [{James Bond 32 [Shaken, not stirred Youth is no guarantee of innovation In his majesty's royal service]} {Miss Moneypenny 27 [James, it is soo good to see you Would you like me to take care of that for you, James? I would really prefer to be a secret agent myself.]} {M Hmmmm 54 [Oh, James. You didn't. Dear God, what has James done now? Can someone please tell me where James Bond is?]}]
	// Person # 0
	// 	 James Bond 32
	// 	         Shaken, not stirred
	// 	         Youth is no guarantee of innovation
	// 	         In his majesty's royal service
	// Person # 1
	// 	 Miss Moneypenny 27
	// 	         James, it is soo good to see you
	// 	         Would you like me to take care of that for you, James?
	// 	         I would really prefer to be a secret agent myself.
	// Person # 2
	// 	 M Hmmmm 54
	// 	         Oh, James. You didn't.
	// 	         Dear God, what has James done now?
	// 	         Can someone please tell me where James Bond is?
}
