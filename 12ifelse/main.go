package main

import "fmt"

func main() {
	fmt.Println("if else in GOLANG")
	loginCount := 11
	var result string

	if loginCount < 10 {
		result = "regular"
	} else if loginCount == 10 {
		result = "proper"
	} else {
		result = "Pro"
	}

	fmt.Println("the user is : ", result)

	if num := 3; num < 10 {
		fmt.Println("number is less than 3")
	} else {
		fmt.Println("number is not less than 3")
	}

	// if err != nill {

	// }

}
