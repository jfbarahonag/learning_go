package main

import (
	"fmt"
	"time"
)

func main() {
	fnc("direct mode")
	go fnc("async mode")

	go func(message string) {
		fmt.Println(message)
	}("Sending message")

	time.Sleep(1 * time.Second)

	fmt.Println("End of routine")

}

func fnc(from string) {
	for i := range 3 {
		fmt.Println(from, ": ", i)
	}
}