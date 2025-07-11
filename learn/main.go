package main

import (
	"fmt"
)

func main() {
	welcome := "This is a program for understanding functions."
	fmt.Println(welcome)

	added := add(35, 34, 50, 50, 34)
	fmt.Println(added)

}

func add(addVar ...int) int {
	totalVar := 0

	for _, val := range addVar{
		totalVar += val
	}

	return totalVar

}