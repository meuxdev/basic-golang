package main

import (
	"fmt"
	pk "main/mypackage"
	"main/pc"
)

func modGlobalValue(val *int) {
	*val = 10
}

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Supra"
	myCar.Year = 2001
	myCar.Model = "MK4"

	fmt.Println(myCar)

	pk.Print("My personal Print from another package")

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
}
