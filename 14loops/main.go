package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")
	days := []string{"Sunday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	// for d:=0; d<len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i:= range days{
	// 	fmt.Println(days[i])
	// }

	// for index, value
	// for index, day := range days{
	// 	fmt.Printf("Index of %v and Value is %v\n", index, day)
	// }

	rougueValue := 1
	for rougueValue < 10 {

		if rougueValue == 2 {
			goto visit
		}

		if rougueValue == 5{
			rougueValue++
			continue

			// break
		}
		fmt.Println("Value is:", rougueValue)
		rougueValue++
	}

	visit:
		fmt.Println("Jumping to visit statement call")

}