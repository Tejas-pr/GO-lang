package main

import "fmt"

func main() {
	// no inheritance in GOLANG
	// no super, parent and child
	hitesh := User{"hitesh", "hitest@gmail.com", true, 16}
	fmt.Println("the user is : ", hitesh)
	fmt.Printf("hitesh details are : %+v\n", hitesh)
	fmt.Printf("the name is %v and the email is %v\n", hitesh.Name, hitesh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
