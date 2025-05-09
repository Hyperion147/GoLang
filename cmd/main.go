package main

import "fmt"

func main(){
	var i = make(chan int)
	go printingChannel(i)
	fmt.Println(<-i)
}

func printingChannel(i chan int){
	i <- 110
}