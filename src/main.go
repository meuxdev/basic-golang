package main

import (
	"fmt"
	"main/pc"
	"math"
)

type figures2D interface {
	area() float64
}

type Square struct {
	base float64
}

type Rectangle struct {
	base   float64
	height float64
}

func (s Square) area() float64 {
	return math.Pow(s.base, 2)
}

func (r Rectangle) area() float64 {
	return r.base * r.height
}

func calculate(f figures2D) {
	fmt.Println("Area: ", f.area())
}

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

	mySquare := Square{base: 3}
	myRectangle := Rectangle{base: 2, height: 4}
	calculate(myRectangle)
	calculate(mySquare)

	// list of interfaces
	myInterface := []interface{}{"Hello", 21, 3.3}
	fmt.Println(myInterface)

}
