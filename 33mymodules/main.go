package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in GLOLANG")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hey the mod user")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to GOLANG</h1>"))
}

// go mod tidy = expensive operation
// go mod verify = verify all modules - take hash and verify
// go list
// go list all
// go list -m all
// go list -m -versions <packageName>
// go mod why <packageName>
// go mod graph
// go mod edit -go
// go mod edit - module
// go mod vendor
// go run -mod=vendor main.go
