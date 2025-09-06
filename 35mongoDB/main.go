package main

import (
	"fmt"
	"log"
	"mangodbb/router"
	"net/http"
)

func main() {
	fmt.Println("mongoDB")
	fmt.Println("server is starting ...")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("the port is listing on the port 4000")
}
