package main

import "fmt"


func main() {
	fmt.Println("Welcome to struct")
	// No inheritance in golang; No super or parent
	
	mayur := User{"Mayur", "patilmayur.2745@gamil.com", true, 21}

	fmt.Printf("mayur details are: %+v\n", mayur)
	fmt.Printf("Name is %v and email is %v\n", mayur.Name, mayur.Email)
	mayur.GetStatus();
	mayur.NewMail()
	fmt.Printf("Name is %v and email is %v\n", mayur.Name, mayur.Email)
	
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus(){
	fmt.Println("Is user active:", u.Status)
}

// Gonna change the original value
// func (u *User) NewMail()

func (u User) NewMail(){
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is:", u.Email)
}