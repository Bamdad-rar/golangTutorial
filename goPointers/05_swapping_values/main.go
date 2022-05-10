package main

import "fmt"


func swap(a *int, b *int){
	tmp := *a
	*a = *b
	*b = tmp
}


func main() {
	x := 10
	y := 20
	fmt.Printf("before (x,y)=(%d,%d)",x,y)

	swap(&x,&y)

	fmt.Printf("after (x,y)=(%d,%d)",x,y)
}
