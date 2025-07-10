package main

import (
	"fmt"
)

func main() {
	welcome := "This is a program for understanding pointers"
	fmt.Println(welcome)
	
	var ptr *int;
	fmt.Println("Value of pointer is", ptr)

	myPtr := 34.5
	var ptr2 = &myPtr
	fmt.Println("Value of 2nd pointer is", ptr2)
	fmt.Println("Value of actual 2nd pointer is", *ptr2)

	*ptr2 = *ptr2 * 2
	fmt.Println("New value of pointer is", *ptr2)

}
