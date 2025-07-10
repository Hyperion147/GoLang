package main

import (
	"fmt"
	"sort"
)

func main() {
	welcome := "This is a program for understanding array and slices"
	fmt.Println(welcome)
	
	var carList [4]string
	
	carList[0] = "Honda" 
	carList[1] = "Toyota"
	carList[2] = "Suzuki"
	carList[3] = "Mahindra"

	fmt.Println("Car company list is:", carList)
	fmt.Println("Car company list is:", len(carList))

	var cars = []string{"Zonda", "Pagani", "Gamera"}
	fmt.Printf("%T \n", cars)
	fmt.Println("Cars are:", cars)

	cars = append(cars, "Maserati", "Regera")
	fmt.Println("Cars are:", cars)

	cars = append(cars[3:4])
	fmt.Println("Cars are:", cars)

	highs := make([]int, 4)

	highs[0] = 5
	highs[1] = 4
	highs[2] = 3
	highs[3] = 2
	fmt.Println("Highs", highs)

	highs = append(highs, 1, 10, 8)
	fmt.Println("Highs", highs)
	fmt.Println(sort.IntsAreSorted(highs))
	
	sort.Ints(highs)
	fmt.Println("Sorted highs", highs)
	fmt.Println(sort.IntsAreSorted(highs))

}
