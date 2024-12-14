// Defer perform LIFO, as line by line exceution is done. The last defer statement will be executed first and prvious ones along.
package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
}

// Hello, 0, 1, 2, 3, 4, two, one, world

func myDefer(){
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}