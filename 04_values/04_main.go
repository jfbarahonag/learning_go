package main

import (
	"fmt"
	"strconv"
)

func main() {
	number1, number2 := 10, 20
	fmt.Println(number1, number2)

	intNumber := 23
	floatNumber := 3.14
	intResult := intNumber + int(floatNumber)
	floatResult := float64(intNumber) + floatNumber
	fmt.Println(intResult, floatResult)

	firstName := "John"
	lastName := "Doe"
	fullName := firstName + " " + lastName
	fmt.Println(fullName)

	// Challenge: Name and age
	name := "Alice"
	age := 30
	intro := "My name is " + name + " and I am " + fmt.Sprint(age) + " years old."
	fmt.Println(intro)
	fmt.Printf("My name is %s and I am %d years old.", name, age)
	fmt.Println("My name is " + name + " and I am " + strconv.Itoa(age) + " years old.")
}
