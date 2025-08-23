package main

import "fmt"

func main() {
	// no inheritance in GOLANG
	// no super, parent and child
	hitesh := User{"hitesh", "hitest@gmail.com", true, 16}
	// fmt.Println("the user is : ", hitesh)
	// fmt.Printf("hitesh details are : %+v\n", hitesh)
	// fmt.Printf("the name is %v and the email is %v\n", hitesh.Name, hitesh.Email)
	hitesh.getStatus()
	hitesh.NewMail()
	fmt.Printf("hitesh details are : %+v\n", hitesh)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) getStatus() {
	fmt.Println("the user status is : ", u.Status)
}

// here it will have the copy of the instance not the original value or the pointer
func (u User) NewMail() {
	u.Email = "test@go.in"
	fmt.Println("Email user is : ", u.Email)
}
