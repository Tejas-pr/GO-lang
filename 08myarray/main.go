package main

import "fmt"

func main() {
	var friutList [4]string
	friutList[0] = "Apple"
	friutList[1] = "Banana"
	friutList[2] = "Orange"
	friutList[3] = "Pinapple"

	fmt.Println("the funn array is ", friutList)
	fmt.Println("the funn array is ", len(friutList)) // this will always measure length that you init

	var vegList = [5]string{"potato", "onion", "beans"}
	fmt.Println("the veg list is : ", vegList)
	fmt.Println("the veg list is : ", len(vegList))
}
