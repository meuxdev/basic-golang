package main

import (
	"fmt"
	"log"
	"strconv"
)

func evenOrOdd() (odd, even bool) {
	// chaleng odd or even
	var numberString string = ""
	var err any
	var val int

	for {
		fmt.Println("Enter Your Number: ")
		fmt.Scanln(&numberString)
		val, err = strconv.Atoi(numberString)
		if err == nil {
			break
		}
	}

	if val%2 == 0 {
		fmt.Println("Is odd")
		return true, false
	} else {
		fmt.Println("Is even")
		return false, true
	}
}

func passwordUser() {
	// challenge to compare password and username
	var inPass string
	var inUser string
	dbUser := "admin"
	dbPass := "admin"

	fmt.Println("Insert Username: ")
	fmt.Scanln(&inUser)
	fmt.Println("Insert Password")
	fmt.Scanln(&inPass)

	if inUser == dbUser && dbPass == inPass {
		fmt.Println("Accessing to your account")
	} else {
		fmt.Println("Invalid credencials")
	}

}

func main() {
	val1 := 1
	val2 := 2

	if val1 == 1 {
		fmt.Println("Value is 1")
	} else {
		fmt.Println("Value is NOT 1")
	}

	if val1 == 1 && val2 == 2 {
		fmt.Println("Is true two conditions are true")
	} else {
		fmt.Println("One of the 2 conditions is NOT true")
	}

	if val1 == 0 || val2 == 2 {
		fmt.Println("On of the two conditions is true")
	}

	// convert text to number
	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value: ", value)

	// challenge odd or even

	evenOrOdd()
	passwordUser()

	// SWITCH
	// first part creates the var ; evaluates the var
	// module := 4 % 2 the var can also be declared outside the switch
	switch module := 4 % 2; module {
	case 0:
		fmt.Println("Is odd")
	default:
		fmt.Println("Is even")
	}

	valueSwitch := 80
	switch {
	case valueSwitch > 100:
		fmt.Println("Is bigger than 100")
	case valueSwitch < 0:
		fmt.Println("Es less than 0")
	default:
		fmt.Println("less than 100 and bigger than 0")
	}
}
