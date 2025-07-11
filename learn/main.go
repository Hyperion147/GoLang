package main

import (
	"fmt"
	
)

func main() {
	welcome := "This is a program for understanding map."
	fmt.Println(welcome)
	
	people := make(map[string]string)

	people["First"] = "First"
	people["Second"] = "First2"
	people["Third"] = "First3"

	fmt.Println("People are:", people)
	
	delete(people, "Second")
	
	fmt.Println("People are:", people)

	for _, personality := range people {
		fmt.Printf("For person , for peosonality %v \n", personality)
	}

}
