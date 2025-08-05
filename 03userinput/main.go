package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the Go programming world!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for the pizza:")
	// commma ok || error ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("thank you for rating, ", input)
	fmt.Printf("the type of the input is %T", input)
}
