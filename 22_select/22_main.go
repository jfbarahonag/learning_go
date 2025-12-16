package main

import (
	"fmt"
	"time"
	"strings"
)

func main() {
	// main1()
	main2()
}

func main1() {
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

func main2() {
	var remainingTickets = (int8)(5)

	ch1 := make(chan int8)
	ch2 := make(chan int8)
	ch3 := make(chan int8)

	go buyTickets(2, &remainingTickets, ch1)
	go buyTickets(1, &remainingTickets, ch2)
	go buyTickets(4, &remainingTickets, ch3)

	for range(3) {
		select {
		case rem := <-ch1:
			fmt.Println("Remaining", rem)
		case rem := <-ch2:
			fmt.Println("Remaining", rem)
		case rem := <-ch3:
			fmt.Println("Remaining", rem)
		}
	}

}

func buyTickets(n int8, remaining *int8, ch chan int8) {
	fmt.Println("Purchasing BF", n)
	time.Sleep(1 * time.Second)
	fmt.Println("Purchasing AS ", n)
	if (*remaining - n) < 0 {
		fmt.Println("Out of stock", *remaining)
		ch <- *remaining * -1
		return
	}
	*remaining -= n
	ch <- *remaining
	fmt.Println("Purchase successfully")
}