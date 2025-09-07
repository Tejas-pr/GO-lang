package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("channels in GOLANG")

	// this will create a deadlock
	// myCh := make(chan int, 1)
	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// myCh <- 5
	// fmt.Println(myCh)
	wg.Add(2)
	// R only
	func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChanelOpen := <-myCh
		fmt.Println(isChanelOpen)
		fmt.Println(val)
		fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)

	// send only
	func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 5
		myCh <- 6
		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
