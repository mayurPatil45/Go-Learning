package main

import "fmt"

const LoginToken = "ibneribi" // PUBLIC

func main() {
	var username string = "Mayur"
	fmt.Println(username)
	fmt.Printf("Variable is of Type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of Type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of Type: %T \n", smallVal)

	var smallFloat float64 = 255.754573895385754854378
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of Type: %T \n", smallFloat)
	
	// Default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of Type: %T \n", anotherVariable)

	// Implicit type
	var website = "google.com"
	fmt.Println(website)

	// No var style
	numberOfUser := 300
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
}