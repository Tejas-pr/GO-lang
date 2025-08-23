package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("files handling")
	context := "I need to put this in the file / text file in the current dir"

	file, err := os.Create("./myfirstFile.txt")
	checkNilErr(err)
	length, err := io.WriteString(file, context)
	checkNilErr(err)
	fmt.Println("the file is : ", length)
	defer file.Close()
	readFile("./myfirstFile.txt")
}

func readFile(fileName string) {
	dataByte, err := os.ReadFile(fileName)
	checkNilErr(err)
	fmt.Println("the file read : ", string(dataByte))
}

// to avoid repeating error check
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
