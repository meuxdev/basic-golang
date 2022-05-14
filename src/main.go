package main

import (
	"fmt"
	pk "main/mypackage"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Supra"
	myCar.Year = 2001
	myCar.Model = "MK4"

	fmt.Println(myCar)

	pk.Print("My personal Print from another package")

}
