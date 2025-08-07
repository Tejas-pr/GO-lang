package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"apple", "banana", "peach"}
	fmt.Printf("the type of this is : %T\n", fruitList)
	fruitList = append(fruitList, "Mongo", "elppa")
	fmt.Println("the new fruit list is : ", fruitList)

	// to slice the slice
	// fruitList = append(fruitList[1:])
	// fmt.Println("the sliced list is : ", fruitList)

	fruitList = append(fruitList[2:5]) // last range are excluded
	fmt.Println("the sliced list is : ", fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 244
	highScores[2] = 212
	highScores[3] = 897
	// highScores[4] = 897 // if you do this it will through error
	// if you use the append method then it will reallocate the memory that we have already create with the size of 4

	highScores = append(highScores, 343, 222, 333)
	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println("the ascending order : ", highScores)
	fmt.Println("is sorted: ", sort.IntsAreSorted(highScores))
}
