package main

import (
	"fmt"
	"math"
)

func sayHello(name string) {
	fmt.Printf("Hello %s, How are you?\n", name)
}

func sayHelloWithAge(name string, age uint32, height, weight float64) {
	fmt.Printf("Hello %s, \nYou have %d years\nYour height is %gcm\nYour weight is %gkg\n", name, age, height, weight)
	// %g -> float32, float64, complex etc
	// %d -> decimal
	// %s  -> string
}

func double(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func pckgExampleFmt() {
	helloMsg := "Hello"
	worldMsg := "World"

	// print ln
	fmt.Println(helloMsg, worldMsg)
	fmt.Println(helloMsg, worldMsg)

	// print F
	name := "Alejandro"
	age := 24
	fmt.Printf("%s has %d years old\n", name, age)
	// %s -> for string
	// %d -> for integer
	// %v -> for any value
	// %g -> for float32 complex etc
	fmt.Printf("%v has %v years old\n", name, age)

	// Springf
	var msg string = fmt.Sprintf("%s has %d years old\n", name, age) // saves this string to var string
	fmt.Println(msg)

	// type data
	alive := true
	fmt.Printf("Type of is alive: %T \n", alive) // %T -> type of the var
	fmt.Printf("Type of is age: %T \n", age)     // %T -> type of the var
}

func circleArea(radius float64) (result float64, errors bool) {
	const pi float64 = 3.1416
	area := pi * math.Pow(radius, 2)
	fmt.Println("*************\nCircle Area\n*************\nRadius: ", radius)
	fmt.Println("Area: ", area)
	return area, false
}

func rectangleArea(sideX, sideY float64) (result float64, err bool) {
	area := sideY * sideX
	fmt.Println("*************\nRectangle Area\n*************\nSide X: ", sideX)
	fmt.Println("Side Y: ", sideY)
	fmt.Println("Result: ", area)
	return area, false
}

func trapezoidArea(top, bottom, height float64) (result float64, err bool) {
	area := (top + bottom) * height / 2
	fmt.Println("*************\n Trapezoid Area\n*************\n Top: ", top)
	fmt.Println("Bottom: ", bottom)
	fmt.Println("Height: ", height)
	fmt.Println("Result: ", area)
	return area, false
}

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

	x := 10
	y := 50

	result := x + y
	fmt.Println("Sum result:", result)

	// diff

	result = y - x
	fmt.Println("Diff result: ", result)

	// times
	result = x * y
	fmt.Println("Times result: ", result)

	// division
	result = y / x
	fmt.Println("Division result:", result)

	// residuo/ module
	result = y % x
	fmt.Println("Module: ", result)

	// Increment
	x++
	fmt.Println("Increment: ", x)
	// decrement
	x--
	fmt.Println("Decrement: ", y)

	// Challenge Area of rectangle, trapeze and circle
	resultAreaCircle, _ := circleArea(3)
	resultAreaRectangle, errors := rectangleArea(12.3, 14.3)
	trapezoidArea(12.3, 14.3, 12.3)
	pckgExampleFmt()
	sayHelloWithAge("Alejando", 24, 1.73, 74)

	val1, val2 := doubleReturn(2)
	fmt.Println("Value1 and value2 ", val1, val2)

	// Escape for the value _ -> in case you dont need of the values that returns
	_, double := doubleReturn(5)
	fmt.Println("Double", double)
	fmt.Println("Result Area Circle: ", resultAreaCircle)
	fmt.Printf("Result Area Rectangle: %g\nErrors: %t ", resultAreaRectangle, errors)

}
