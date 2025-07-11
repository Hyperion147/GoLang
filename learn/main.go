package main

import (
	"fmt"
)

func main() {
	welcome := "This is a program for understanding structs."
	fmt.Println(welcome)

	cs2 := Game{"cs2", "pc", 10, true}
	fmt.Println(cs2)
	fmt.Printf("CS2 player count: %+v\n", cs2.Players)
	fmt.Printf("CS2 details: %+v", cs2)

}

type Game struct {
	Name     string
	Platform string
	Players  int
	Status   bool
}
