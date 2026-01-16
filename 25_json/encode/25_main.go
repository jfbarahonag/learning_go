package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
}

func main() {
	person := Person {
		Name: "Felipe",
		Age: 29,
		Email: "felipe@mail.com",
	}

	jsonData, err := json.Marshal(person)

	if err != nil {
		fmt.Println("Error encoding to JSON", err)
		return
	}

	fmt.Println("JSON data: ", string(jsonData))
}