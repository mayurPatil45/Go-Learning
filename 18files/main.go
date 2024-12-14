package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a files = www.google.com"

	file, err := os.Create("C:/Users/hp/Documents/GoLearning/18files/test.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println("Length is:",length)
	defer file.Close()
	readFile("./test.txt")
}

func readFile(filename string){
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)
	fmt.Println("Text data inside the file is \n", string(databyte))
}

func checkNilError(err error){
	if err != nil {
		panic(err)
	}
}