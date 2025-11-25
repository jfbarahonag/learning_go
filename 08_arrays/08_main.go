
package main

import (
	"fmt"
)

func main() {
	
	var arr [5]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30

	fmt.Println("Array values:", arr)
	fmt.Println("Array length:", len(arr))

	separator()

	ages := [3]int{25, 30, 35}
	fmt.Println("Ages array:", ages)

	separator()

	agesCopy := ages
	agesCopy[0] = 40

	fmt.Println("Original ages array:", ages)
	fmt.Println("Copied ages array:", agesCopy)

	separator()

	numbers := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println("Numbers array:", numbers)
	fmt.Println("Numbers array length:", len(numbers))
}

func separator() {
	fmt.Println("------------------------------")
}
