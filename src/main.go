package main

import (
	"fmt"
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

func main() {
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
