package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to maps")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["Py"] = "Python"

	fmt.Println("List of all the languges", languages)

	// IF we want to get the value, provide the key
	fmt.Println("JS shorts for", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all the languges", languages)

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v. \n", key, value)
	}
}