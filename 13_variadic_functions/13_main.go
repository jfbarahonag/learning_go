package main

import (
	"fmt"
)

func main() {
	average := avg(10.5, 20.0, 30.5, 40.0)
	fmt.Printf("The average is: %.2f\n", average)

	values := []float64{5.0, 15.0}
	average = avg(values...)
	fmt.Printf("The average is: %.2f\n", average)

	average = avg()
	fmt.Printf("The average is: %.2f\n", average)
}

func avg(numbers ...float64) float64 {

	length := len(numbers)

	if length == 0 {
		return 0.0
	}

	total := 0.0

	for _, number := range numbers {
		total += number
	}

	return total / float64(length)

}