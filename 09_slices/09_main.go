package main

import (
	"fmt"
	"slices"
)

func main() {
	//stringArr := []string{}
	var stringArr []string

	fmt.Println("data:", "content",stringArr, 
	"condition:", stringArr == nil, "length:", len(stringArr) == 0)

	stringSlice := make([]string, 3)
	stringSlice[0] = "Hello"
	stringSlice[1] = "from"
	stringSlice[2] = "Go!"
	//stringSlice[4] = "Extra" // Panic: index out of range
	fmt.Println("data:", stringSlice)
	fmt.Println("Specific data:", stringSlice[2])
	fmt.Println("length:", len(stringSlice))

	stringSlice = append(stringSlice, "Extra1", "Extra2")
	fmt.Println("data after append:", stringSlice)
	fmt.Println("length after append:", len(stringSlice))

	secondArray := []string{"One", "Two", "Three"}
	thirdArray := []string{"One", "Two", "Three"}

	if slices.Equal(secondArray, thirdArray) {
		fmt.Println("The slices are equal")
	} else {
		fmt.Println("The slices are NOT equal")
	}

	stringSlice = append(stringSlice, secondArray...)
	fmt.Println("data after second append:", stringSlice)
	fmt.Println("length after second append:", len(stringSlice))
}

// func separator() {
// 	fmt.Println("------------------------------")
// }
