package main

import (
	"fmt"
	"strings"
)


func main() {
	
	i := 0
	for i < 3 {
		fmt.Printf("Iteration %d\n", i)
		i++
	}
	
	separator()
	
	for idx := 0; idx < 3; idx++ {
		fmt.Printf("Iteration %d\n", idx)
	}
	
	separator()
	
	for idx := range 3 {
		fmt.Printf("Iteration %d\n", idx)
	}

	separator()

	for {
		fmt.Println("Loop")
		break
	}

	separator()

	for value := range 6 {
		if value%2 == 0 {
			continue
		}
		fmt.Println(value)
	}

	separator()

	names := []string{"Ana", "Luis", "Carlos"}
	names = append(names, "Maria", "Pepe")

	for idx, name := range names {
		fmt.Println(idx + 1, strings.ToUpper(name))
	}

}

func separator() {
	fmt.Println(strings.Repeat("-", 30))
}
