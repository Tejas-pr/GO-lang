package main

import "fmt"

func main() {
	fmt.Println("Welcome class on pointers")

	// var ptr *int
	// fmt.Println("the default value is: ", ptr)

	// refrence means => &
	myNumber := 56
	var ptr = &myNumber
	// pointer is refrence to the direct memory location
	fmt.Println("the value of actual pointer is : ", ptr)
	fmt.Println("the value of actual pointer is : ", *ptr)

	// the actual value
	*ptr = *ptr + 2
	fmt.Println("the final answer is : ", myNumber)
	// here we are directly adding the value to the exact value from the address
	// this will garenty that opeartion will performed on the specified value.
}
