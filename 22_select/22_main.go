package main

import (
	"fmt"
	"time"
	"strings"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)
	channel3 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		channel1 <- "first message"
	}()
	
	go func() {
		time.Sleep(3 * time.Second)
		channel2 <- "second message"
	}()
	
	go func() {
		time.Sleep(2 * time.Second)
		channel3 <- "third message"
	}()

	var builder strings.Builder
	for range 3 {
		select {
		case msg := <-channel1:
			builder.WriteString(msg+" | ")
			fmt.Println("Channel 1: ", msg)
		case msg := <-channel2:
			builder.WriteString(msg+" | ")
			fmt.Println("Channel 2: ", msg)
		case msg := <-channel3:
			builder.WriteString(msg+" | ")
			fmt.Println("Channel 3: ", msg)
		}
	}
	finalMsg := builder.String()
	fmt.Println("FINISH", finalMsg)
}