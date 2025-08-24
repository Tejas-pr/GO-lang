package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"cousename"`
	Price    int
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	EncodeJson()
}

func EncodeJson() {
	lcoCourse := []course{
		{"reactjs", 299, "abcd", []string{"web-dev", "JS"}},
		{"nextjs", 199, "abcdef", []string{"full-dev", "TS"}},
		{"angular", 199, "abcdefgh", nil},
	}
	// package this to JSON data
	// finalJSON, err := json.Marshal(lcoCourse)
	finalJSON, err := json.MarshalIndent(lcoCourse, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJSON)
}
