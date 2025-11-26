package main

import (
	"fmt"
)

func main() {
	fmt.Println("1 + 2 = ", sum(1,2))
	fmt.Println("5 - 3 = ", subtract(5,3))
	fmt.Println("4 * 6 = ", multiply(4,6))
	quotient, remainder := divide(10, 3)
	fmt.Printf("10 / 3 = %d with a remainder of %d\n", quotient, remainder)
	quotient, remainder = divide(10, 0)
	if quotient == 0 && remainder == 0 {
		fmt.Println("Cannot divide by zero")
	} else {
		fmt.Printf("10 / 0 = %d with a remainder of %d\n", quotient, remainder)
	}

	num1 := getNumber()
	num2 := getNumber()
	num3 := getNumber()

	sumResult := sumAll(num1, num2, num3)
	restResult := substractAll(num1, num2, num3)

	fmt.Println("The sum of all the numbers is:", sumResult)
	fmt.Println("The rest of all the numbers is:", restResult)
}

func getNumber() int {
	var number int
	fmt.Println("Type a number:")
	fmt.Scanln(&number)
	return number
}

func sum(num1, num2 int) int {
	return num1 + num2
}

func sumAll(numbers ...int) int {
	total := numbers[0]

	for _, number := range numbers[1:] {
		total += number
	}

	return total	
}

func subtract(num1, num2 int) int {
	return num1 - num2
}

func substractAll(numbers ...int) int {
	total := numbers[0]
	
	for _, number := range numbers[1:] {
		total -= number
	}

	return total
}

func multiply(num1, num2 int) int {
	return num1 * num2
}

func divide(num1, num2 int) (int, int) {

	if num2 == 0 {
		return 0, 0
	}

	quotient := num1 / num2
	remainder := num1 % num2
	return quotient, remainder
}