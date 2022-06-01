package main

import "fmt"

// getting a pointer as input
func testFunc(myPointer *int){

	// dereferencing the pointer, which gives us access to the variable that is holding
	// an integer value and we change the value to 420
	*myPointer = 420
}


func main() {
	// creating a pointer
	aPointer := new(int)

	fmt.Println(*aPointer)
	testFunc(aPointer)
	// print the value by dereferncing and gaining access to it.
	fmt.Println(*aPointer)
}
