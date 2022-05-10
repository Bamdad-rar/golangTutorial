package main

import "fmt"

/*
pointers reference where the location/address of a value is.
rather than the value itself.

in go a pointer is represented by * followed by the type of the stored value
so a pointer to an integer is : *int

* is also used in dereferencing. dereferencing lets us access the value that
the pointer points at.

& is used to find the "address/location" of the "variable".
so &x for example stores the pointer to where the desired variable is.

|a		|	|	|b	|	|	|	|
|*int:3	|	|	|10	|	|	|	|

*/

func main() {
	x := 10
	
	var aPointer *int = &x
	
	fmt.Println("value of x:",x)
	fmt.Println("value of aPointer:",aPointer)
	fmt.Println("x address:",&x)
	fmt.Println("value of x via pointer dereferencing:",*aPointer)
	
}