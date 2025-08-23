package main

import (
	"fmt"
)

func main() {
	fmt.Println("loops =====")

	// days := []string{"monday", "tuesday"}

	// for index, day := range days {
	// 	fmt.Printf("the key is %v and value is %v\n", index, day)
	// }

	// toRange := 1

	// for toRange <= 10 {
	// 	fmt.Println("the value is :", toRange)
	// 	toRange++
	// }

	result := addSum(3, 5)
	fmt.Printf("the value of the number is %v\n", result)

	result, myMessage := proAdder(4, 5, 6, 7, 8, 9)
	fmt.Println("the pro sum is : ", result, myMessage)
}

func addSum(val1 int, val2 int) int {
	return val1 + val2
}

func proAdder(value ...int) (int, string) {
	total := 0
	for _, val := range value {
		total += val
	}
	return total, "tejas"
}
