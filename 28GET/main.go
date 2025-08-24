package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// PerfornGET()
	// PerfornPOST()
	PerfornPOSTFormRequest()
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

func PerfornPOST() {
	const myurl = "http://localhost:3000/post"

	// JSON payload
	reqestBody := strings.NewReader(`
		{
			"courseName": "fullstack",
			"name": "tejas"
		}
	`)
	res, err := http.Post(myurl, "application/json", reqestBody)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))
}

func PerfornPOSTFormRequest() {
	const myurl = "http://localhost:3000/post"

	//form data
	data := url.Values{}
	data.Add("firstName", "tejas")
	data.Add("lastName", "PR")
	data.Add("email", "tejas@gmail.com")

	res, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	content, _ := io.ReadAll(res.Body)
	defer res.Body.Close()
	fmt.Println(string(content))
}
