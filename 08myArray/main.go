package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string 

	fruitList[0] = "Apple"
	fruitList[1] = "Guava"
	fruitList[3] = "Pinapple"

	fmt.Println("Fruit List is", fruitList)
	fmt.Println("Fruit List is", len(fruitList))

	var vegList = [5]string{"Potato", "Tomato", "Beans"}
	fmt.Println("Vegy list is", vegList)
	fmt.Println("Vegy list is", len(vegList))
}