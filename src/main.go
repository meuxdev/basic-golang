package main

import "fmt"

type Car struct {
	brand string
	model string
	power uint
	year  uint
}

func main() {

	myCar := Car{brand: "Supra", model: "mk4"}
	fmt.Println(myCar)

	// other way to instance

	var otherCar Car
	otherCar.brand = "Ferrari"

	fmt.Println(otherCar)
}
