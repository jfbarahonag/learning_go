package main

import (
	"fmt"
	"maps"
)

func main() {
	myMap := make(map[string]int)
	myMap["Felipe"] = 29
	myMap["Vivi"] = 27

	fmt.Println("My map: ", myMap)

	felipe := myMap["Felipe"]
	fmt.Println("Felipe's age:", felipe)

	vivi := myMap["Vivi"]
	fmt.Println("Vivi's age:", vivi)
	
	_, exists := myMap["Pepe"]
	fmt.Println("Does Pepe exist in the map?", exists)

	fmt.Println("NNumber of elements in the map:", len(myMap))

	delete(myMap, "Vivi")
	fmt.Println("My map after deleting Vivi:", myMap)

	delete(myMap, "Andres")
	fmt.Println("My map after trying to delete Andres:", myMap)

	clear(myMap)
	fmt.Println("My map after clearing all elements:", myMap)
	
	separator()

	newMap1 := map[string]string{
		"USA":    "Washington D.C.",
		"France": "Paris",
		"Japan":  "Tokyo",
	}
	
	newMap2 := map[string]string{
		"USA":    "Washington D.C.",
		"France": "Paris",
		"Japan":  "Tokyo",
	}

	if maps.Equal(newMap1, newMap2) {
		fmt.Println("The maps are equal")
	}
	
}

func separator() {
	fmt.Println("------------------------------")
}