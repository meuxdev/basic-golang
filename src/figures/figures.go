package figures

import (
	"fmt"
	"math"
)

// Struct Rectangle
type Rectangle struct {
	Base   float64
	Height float64
}

func (r Rectangle) area() float64 {
	return r.Base * r.Height
}

// Struct for triangle
type Triange struct {
	Base   float64
	Height float64
}

func (t Triange) area() float64 {
	return (t.Base * t.Height) / 2
}

// Struct Square
type Square struct {
	Side float64
}

func (s Square) area() float64 {
	return math.Pow(s.Side, 2)
}

type figures2D interface {
	area() float64
}

// Prints the area of a ficure 2D
func CalculateArea(f figures2D) {
	fmt.Println("Area: ", f.area())
}
