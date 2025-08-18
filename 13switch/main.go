package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("swich case in GOLANG")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("the value of the dice is : ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Printf("The dice value is %d and you can start\n", diceNumber)
	case 2:
		fmt.Printf("in spot 2")
	case 3:
		fmt.Printf("in spot 3")
	case 4:
		fmt.Printf("in spot 4")
		// fallthrough
	case 5:
		fmt.Printf("in spot 5")
	case 6:
		fmt.Printf("in spot 6 and roll again")
	}

}
