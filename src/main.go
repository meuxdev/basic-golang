package main

import (
	"fmt"
	"main/pc"

	"main/figures"
)

func modGlobalValue(val *int) {
	*val = 10
}

func main() {

	randomNumber := 30
	pointer := &randomNumber // pointer

	fmt.Println(*pointer, pointer)
	fmt.Printf("%T\n", pointer)

	modGlobalValue(pointer)
	fmt.Println(randomNumber)

	myPc := pc.Pc{Ram: 16, Disk: 200, Brand: "MSI"}
	myPc.Ping()

	fmt.Println(myPc)

	myPc.DuplicateRam()

	fmt.Println(myPc)
	// list of interfaces
	myInterface := []interface{}{"Hello", 21, 3.3}
	fmt.Println(myInterface)

	// Figures
	rectangle := figures.Rectangle{Base: 2, Height: 21}
	triangle := figures.Triange{Base: 2, Height: 10}
	square := figures.Square{Side: 20}
	figures.CalculateArea(rectangle)
	figures.CalculateArea(triangle)
	figures.CalculateArea(square)

}
