package main

import(
	"fmt"
	"errors"
)

func main() {
	var value string = "Hello another function"
	anotherFunction(value)

	var num int = 10
	var den int = 0
	var result, err = intDivision(num, den)
	switch result{
	case 0:
		 fmt.Println(err.Error())
	default: 
			fmt.Printf("\nResult is : %v",result)
	}
}

func anotherFunction(value string) {
	fmt.Println(value)
}

func intDivision(num int, den int) (int, error) {
	var err error
	if den == 0 {
		err = errors.New("\nError : Denominator Is Zero")
		return 0, err
	}
	var result int = num/den
	return result, err
}