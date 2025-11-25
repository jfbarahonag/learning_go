
package main

import (
	"fmt"
	"time"
)

func main() {
	
	input := 2
	fmt.Printf("Type %d as ", input)

	switch input {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Unknown number")
	}

	today := time.Now()
	fmt.Println(today)

	switch weekday := today.Weekday(); weekday {
	
	case time.Saturday, time.Sunday:
		fmt.Println(weekday, "It's the weekend")
	default:
		fmt.Println(weekday, "It's a weekday")
	}

	switch {
		case today.Hour() < 12:
			fmt.Println("Good morning")
		case today.Hour() < 18:
			fmt.Println("Good afternoon")
		default:
			fmt.Println("Good evening")
	}

	switch today.Weekday() {
	case time.Monday:
		fmt.Println("Lunes")
	case time.Tuesday:
		fmt.Println("Martes")
	case time.Wednesday:
		fmt.Println("Miércoles")
	case time.Thursday:
		fmt.Println("Jueves")
	case time.Friday:
		fmt.Println("Viernes")
	case time.Saturday:
		fmt.Println("Sábado")
	case time.Sunday:
		fmt.Println("Domingo")
	default:
		fmt.Println("Día inválido")
	}
}
