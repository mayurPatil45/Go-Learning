package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println("Welcome to Slices")

	var fruitList = []string{"Apples", "Tomato", "Peach"}
	// fmt.Printf("Tye of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	// fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	// fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 876
	// highScores[4] = 777

	highScores = append(highScores, 555, 666, 777)

	// fmt.Println(highScores)
	
	// fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	// fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))

	// How to remove a value from slices based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
