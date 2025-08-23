package main

import "fmt"

func main() {
	defer fmt.Println("defer line")
	defer fmt.Println("Two")
	defer fmt.Println("one")
	fmt.Println("Hello")
	// the code will run the LIFO fashion
	// the defer code will execute at the last and which ever come last will put out first
	myDefer()
	// hello, 4321, one, Two, defer line
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
