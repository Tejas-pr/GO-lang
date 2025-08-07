package main

import (
	"fmt"
	"time"
)

func main() {
	getTime := time.Now()
	fmt.Println(getTime)

	fmt.Println(getTime.Format("01-02-2006 15:04:05 Monday"))

	// to create my own date
	createDate := time.Date(2020, time.August, 4, 0, 0, 0, 0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 15:04:05 Monday"))
	// go through go docs - time
	// use go build - to create a executable file from GO
	// if you want to create for window or linex use flag infront of the go build => GOOD="windows" go build
}
