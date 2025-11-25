package main

import (
	"fmt"
)

func main() {
	name := "Felipe"
	age := 29

	if age >= 30 {
		fmt.Printf("%s, you reached third floor", name)
	} else {
		fmt.Printf("%s, you are behind third floor", name)
	}

	fmt.Println()

	if 8%2 == 0 {
		fmt.Println("8 is an even number")
	} else {
		fmt.Println("8 is an odd number")
	}

	// declaring variable inside if statement
	if number := 8; number > 10 {
		fmt.Printf("%d is greater than 10\n", number)
	} else if number > 5 {
		fmt.Printf("%d is greater than 5 but less than or equal to 10\n", number)
	} else {
		fmt.Printf("%d is less than or equal to 5\n", number)
	}
}
