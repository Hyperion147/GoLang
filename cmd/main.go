package main

import ( "fmt" )

func main(){
	var intSlice []int32 = []int32 {2,3,5}
	var intSlice2 []int32 = []int32 {2,3,5}
	fmt.Println(intSlice)
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 6)
	fmt.Println(intSlice3)
	intSlice3 = append(intSlice3, 7,8,9)
	fmt.Println(intSlice3)
}