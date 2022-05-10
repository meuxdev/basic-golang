package main

import "fmt"

func main() {

	// declarative type
	const pi float64 = 3.14
	// Inferencial
	const pi2 = 3.15

	fmt.Println("Hello World")
	fmt.Print("PI value: ")
	fmt.Println(pi)
	fmt.Println("PI second value", pi2)

	// var integers
	base := 12
	var height int = 14
	var area int

	fmt.Println(base, height, area)

	// zero values
	var zeroValueInteger int
	var zeroValueFloat float64
	var zeroValueString string
	var zeroValueBool bool

	fmt.Println(zeroValueInteger, zeroValueFloat, zeroValueString, zeroValueBool)

	// square area
	const squareBase = 10
	squareArea := squareBase * squareBase
	fmt.Println("Square Area is: ", squareArea)
}
