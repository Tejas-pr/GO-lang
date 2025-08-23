package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "" // this code will read the html

func main() {
	response, err := http.Get(url)
	checkErrNill(err)
	// *http.Response - here you will actual response not the copy of the response
	fmt.Printf("the type of the response is %T\n", response)
	defer response.Body.Close() // dont forget to close the requests

	dataByte, err := io.ReadAll(response.Body)
	checkErrNill(err)
	content := string(dataByte)
	fmt.Println("the data is =====", content)
}

func checkErrNill(err error) {
	if err != nil {
		panic(err)
	}
}
