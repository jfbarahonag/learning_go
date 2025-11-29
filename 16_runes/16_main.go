package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	hello := "안녕하세요"
	hola := "Hola"

	printData(hello)
	printData(hola)
}

func printData(s string) {
	fmt.Println("String:", s)
	fmt.Println("Length in bytes:", len(s))

	for idx := 0; idx < len(s); idx++ {
		fmt.Printf("%x ", s[idx])
	}
	fmt.Println()

	runeCount := utf8.RuneCountInString(s)
	fmt.Println("Number of runes:", runeCount)

	for idx, rune := range s {
		fmt.Printf("Rune %U starts in %d byte\n", rune, idx)
	}
	fmt.Println()
}