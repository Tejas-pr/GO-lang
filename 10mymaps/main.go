package main

import "fmt"

func main() {
	// how to make map in go
	language := make(map[string]string)

	language["JS"] = "Javascript"
	language["RB"] = "Ruby"
	language["PY"] = "Python"

	fmt.Println("List of all languages : ", language)
	fmt.Println("List of all languages : ", language["JS"])

	delete(language, "RB")
	fmt.Println("List of all languages : ", language)

	// loops
	for key, value := range language {
		fmt.Printf("the key is %v and the value is %v\n", key, value)
	}

	for _, value := range language {
		fmt.Printf("the value is %v\n", value)
	}

	language2 := make(map[string]string)

	language2["TJ"] = "tejas"
	language2["TJP"] = "tejasPR"

	for key, value := range language2 {
		fmt.Println("the new language two is : ======================")
		fmt.Printf("the key is %v and the value is %v\n", key, value)
	}
}
