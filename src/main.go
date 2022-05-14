package main

import "fmt"

func main() {
	agesMap := make(map[string]int)

	agesMap["Alejandro"] = 24
	agesMap["Pepito"] = 19

	fmt.Println("Ages Map: ", agesMap)

	// range

	// IMPORTANT: map does not has order to display
	// maps use concurrency
	// more efficient than slices

	for i, value := range agesMap {
		fmt.Println(i, value)
	}

	// find a value

	ageValue, ok := agesMap["Pepito"]
	fmt.Println(ageValue, ok)

	ageValue, ok = agesMap["Josep"]
	fmt.Println(ageValue, ok)
}
