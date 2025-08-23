package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://dummyjson.com/todos:3000?token=sdjsjdsfdjsfls&payment=sdlfjknsdf"

func main() {
	fmt.Println("the url is : ", myUrl)
	// parsing url
	result, _ := url.Parse(myUrl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	// better way to store the query
	qparams := result.Query()
	fmt.Printf("the data type is %T\n", qparams)
	fmt.Printf("the data type is %T\n", qparams["payment"])
	fmt.Println("the better store is: ", qparams)

	for _, val := range qparams {
		fmt.Println("the value of param : ", val)
	}

	// to create a URLs
	// here you will pass the refrence of the URL
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/learn",
		RawQuery: "user=hitesh",
	}

	newUrl := partsOfUrl.String()
	fmt.Println("the new url: ", newUrl)
}
