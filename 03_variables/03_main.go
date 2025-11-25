package main

import "fmt"

func main() {
	var text = "Hello, Variables!"
	fmt.Println(text)

	var secondText string
	secondText = "This is another variable."
	fmt.Println(secondText)

	var intNumber1, intNumber2 = 10, 20

	fmt.Println(intNumber1, intNumber2)

	var booleanVar = true
	fmt.Println(booleanVar)

	var floatNumber float32 = 3.14
	fmt.Println(floatNumber)

	fruit := "Apple"
	fmt.Println(fruit)

	var myVar bool
	fmt.Println(myVar)
	myVar = true
	fmt.Println(myVar)
}
