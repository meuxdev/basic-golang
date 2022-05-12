package main

import "fmt"

func main() {

	// first part creates the var ; evaluates the var
	// module := 4 % 2 the var can also be declared outside the switch
	switch module := 4 % 2; module {
	case 0:
		fmt.Println("Is odd")
	default:
		fmt.Println("Is even")
	}

	value := 80
	switch {
	case value > 100:
		fmt.Println("Is bigger than 100")
	case value < 0:
		fmt.Println("Es less than 0")
	default:
		fmt.Println("less than 100 and bigger than 0")
	}
}
