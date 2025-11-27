package main

import (
	"fmt"
)

func main() {
	value := float64(6)
	fmt.Printf("Factorial of %.0f is: %.0f\n", value, factorial(value))

	fmt.Printf("Fibonacci of %.0f is: %.0f\n", value, fibonacci(value))
}

func factorial(n float64) float64 {
	if n == 0 {
		return 1
	}
	return n * factorial(n - 1)
}

func fibonacci(n float64) float64 {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}