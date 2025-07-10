package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "This is a simple input for getting discord username"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your discord username:")
	input, _ := reader.ReadString('\n')
	fmt.Println("Your username is:", input)
	increment, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The incremented name is:", increment + 1)
	}
}
