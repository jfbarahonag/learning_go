package main

import (
	"fmt"
	"os"
)

func main() {
	envVar := os.Getenv("HOME")
	if envVar == "" {
		fmt.Println("The HOME environment variable is not set.")
	} else {
		fmt.Printf("The HOME environment variable is: %s\n", envVar)
	}

	file, err := os.Create("file.txt")

	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()
}