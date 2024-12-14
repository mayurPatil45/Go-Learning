package main

import "fmt"

func main() {
	fmt.Println("Welcome to function in golang")
	greeter()

	result := adder(3, 2)
	fmt.Println("Result is:", result)
	
	proResult, myMessage := proAdder(2,32 ,5 ,52 ,6)
	fmt.Println("Pro Result is:", proResult)
	fmt.Println("Pro Message is:", myMessage)

}

// If we don't know the number of inputs then use this ...
// func proAdder(values ...int) int {
// 	total := 0
	
// 	for _, val := range values {
// 		total += val
// 	}

// 	return total
// }

func proAdder(values ...int) (int, string) {
	total := 0
	
	for _, val := range values {
		total += val
	}

	return total, "Hi from proResult function"
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func greeter() {
	fmt.Println("Namstey from golang")
}