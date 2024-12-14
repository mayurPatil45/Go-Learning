package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Welcome to struct")
	// No inheritance in golang; No super or parent

	mayur := User{"Mayur", "patilmayur.2745@gamil.com", true, 21}

	fmt.Printf("mayur details are: %+v\n", mayur)
	fmt.Printf("Name is %v and email is %v", mayur.Name, mayur.Email)

}
