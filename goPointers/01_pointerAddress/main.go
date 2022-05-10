package main

import "fmt"


func testFunc(x string){
	fmt.Println("inside testFunc")
	fmt.Println("x memory address:",&x)
}


func main() {
	var name string = "bamdad"
	fmt.Printf("variable 'name'=%s, %p\n",name,&name)

	testFunc(name)

	/* we get different addresses for the same variable, this is because in go
	the varibale that is passed into a function is copied!
	so we're not dealing with the original variable inside the function
	*/
}
