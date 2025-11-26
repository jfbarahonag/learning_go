package main

import (
	"fmt"
)

func main() {
	firstMap, secondMap := multimap("Juan", "Felipe", "Barahona")
	fmt.Println("First Map:", firstMap)
	fmt.Println("Second Map:", secondMap)

	firstPerson := []string{"Juan", "Felipe", "Barahona"}
	secondPerson := []string{"Ana", "Gonzalez"}
	name1, name2 := names(firstPerson, secondPerson)
	fmt.Println("Person 1:", name1, "-")
	fmt.Println("Person 2:", name2, "-")
	fmt.Printf("Person 1: %s-\n", name1)
	fmt.Printf("Person 2: %s-\n", name2)
}

func multimap(firstname, middlename, lastname string) (map[string]string, map[string]string) {
	firstMap := make(map[string]string)
	secondMap := make(map[string]string)
	
	firstMap[firstname] = lastname
	secondMap[middlename] = lastname
	
	return firstMap, secondMap
}

func names(name1, name2 []string) (string, string) {
	names := [][]string{name1, name2}
	returnNames := make([]string, 0)
	
	for _, nameArr := range names {
		var currName string
		for _, name := range nameArr {
			currName += name + " "
		}
		// trim the last space
		currName = currName[:len(currName)-1]
		returnNames = append(returnNames, currName)
	}

	return returnNames[0], returnNames[1]
}