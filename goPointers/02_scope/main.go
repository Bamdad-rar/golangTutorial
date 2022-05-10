package main

import "fmt"

// this wont change the value of x outside of the function's scope
func zero(x int){
	x = 0
}
// one way to bypass the scope is by using a special data type, known as the pointer

func zeroP(xPointer *int){
	*xPointer = 0
}


func main() {
	x := 10
	zero(x)
	fmt.Println(x)

	zeroP(&x)
	fmt.Println(x)
}

