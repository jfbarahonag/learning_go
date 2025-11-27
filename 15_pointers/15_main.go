package main

import (
	"fmt"
	"time"
)

func main() {
	numbers := [3]int{10, 20, 30}

	fmt.Println("Address of numbers", numbers)
	fmt.Printf("Address of numbers %p\n", &numbers)

	timeTaken(func() { updateArray(&numbers) })
	fmt.Println("Updated numbers", numbers)

	timeTaken(func() { updateArrayEfficient(&numbers) })
	fmt.Println("Updated numbers efficient", numbers)
}

// This way works but waste more memory
func updateArray(arr *[3]int) {
	arr[0] = 100
}

// This way is more efficient due to pointer dereferencing
func updateArrayEfficient(arr *[3]int) {
	(*arr)[0] = 200
}

// Function to calculate time taken by another function
func timeTaken(f func()) {
	start := makeTimestamp()
	f()
	end := makeTimestamp()
	fmt.Printf("Time taken: %d ms\n", end-start)
}

func makeTimestamp() int64 {
	return time.Now().UnixNano()// / int64(time.Millisecond)
}
