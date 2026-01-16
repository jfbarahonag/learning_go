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
	//jsonStr := "{\"name\":\"Felipe\",\"age\":29,\"email\":\"felipe@mail.com\"}"
	jsonStr := `{"name":"Felipe","age":29,"email":"felipe@mail.com"}`

	var person Person

	err := json.Unmarshal([]byte(jsonStr), &person)

	if err != nil {
		fmt.Println("Error decoding JSON.", err)
		return
	}

	fmt.Println("Person ->", person.Name, person.Age, person.Email)
}