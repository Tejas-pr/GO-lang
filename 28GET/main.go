package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	PerfornGET()
}

func PerfornGET() {
	const myurl = ""
	res, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println("the status is: ", res.StatusCode)
	fmt.Println("length is : ", res.ContentLength)

	content, _ := io.ReadAll(res.Body)
	fmt.Println("the content is : ", string(content))

	var responseString strings.Builder // this will always hold the response value
	byteCount, _ := responseString.Write(content)
	fmt.Println("the byteCount is : ", byteCount)
	fmt.Println(responseString.String())
}
