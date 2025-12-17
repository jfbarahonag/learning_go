package main

import (
	"fmt"
	"time"
)

func main() {
	counter1 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		counter1 <- "Result 1"
	}()

	select {
	case result := <-counter1:
		fmt.Println("Received", result)
	case <- time.After(1 * time.Second):
		fmt.Println("Timeout reached")
	}
	// --------------------------------
	counter2 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		counter2 <- "Result 1"
	}()

	select {
	case result := <-counter2:
		fmt.Println("Received", result)
	case <- time.After(3 * time.Second):
		fmt.Println("Timeout reached")
	}
}
