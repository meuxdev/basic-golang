package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func isPalindrome(text string) {
	var textReverse string
	// lower case to prevent error of mayus
	text = strings.ToLower(text)

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i]) // we need to parse because the position returns the value in ASCII
	}

	if textReverse == text {
		fmt.Println("Is palindrome")
	} else {
		fmt.Println("Is NOT palindrome")
	}
}

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
		// KEYWORDS CYCLE
		// Defer
		defer fmt.Println("Hello") // last print waits to all to execute
		// defer executes before main is completed.
		// conexion -> defer cierre conexion
		// abrir archivo -> con defer lo cerramos
		// GOD PRACTICE ONLY USE 1 DEFER PER FUNCTION
		fmt.Println("World")

		// continue and break

		for i := 0; i < 10; i++ {
			fmt.Println(i)

			// continue
			if i == 2 {
				fmt.Println("Es 2")
				continue // example: when a error executes but you dont want to break
			}

			// break
			if i == 8 {
				fmt.Println("Break")
				break // when you want to finalize the cycle
			}
		}

		var array [4]int // array inmutables
		array[0] = 1
		array[1] = 2
		fmt.Println(array)
		fmt.Printf("Length of the array is %d\n", len(array))
		fmt.Printf("Max capacity of the array is %d\n", cap(array))

		// slice
		// We dont specify the amount of elemnts that we will get
		slice := []int{0, 1, 2, 3, 4, 5, 6}
		fmt.Println(slice)
		fmt.Printf("Length of the slice is %d\n", len(slice))
		fmt.Printf("Max capacity of the slice is %d\n", cap(slice))

		// slice method
		fmt.Println(slice[0])
		fmt.Println(slice[:3])
		fmt.Println(slice[2:4])
		fmt.Println(slice[4:])

		// slicing[inclusive: exclusive]
		// first number is inclusive and the second one is exclusive

		// slicing method append
		slice = append(slice, 123)
		fmt.Println("new slice: ", slice)

		// appeng new list
		newSlice := []int{8, 9, 10}
		slice = append(slice, newSlice...)
		fmt.Println("New slicing with lists... ", slice)
	}

	slice := []string{}
	slice = append(slice, "Hello")
	slice = append(slice, "how")
	slice = append(slice, "are")
	slice = append(slice, "you?")

	for _, value := range slice {
		fmt.Print(" ", value)
	}

	fmt.Println()

	// for range returns
	// index, value
	for i := range slice {
		fmt.Print(" ", i)
	}

	fmt.Println()

	isPalindrome("ama")
	isPalindrome("aMoR a roma")
	isPalindrome("casas")
}
