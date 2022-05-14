package main

import (
	"fmt"
	"main/pc"
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
}
