package main

import "fmt"

func main() {
	var array [4]int // array inmutables
	array[0] = 1
	array[1] = 2
	fmt.Println(array)
	fmt.Printf("Length of the array is %d\n", len(array))
	fmt.Printf("Max capacity of the array is %d\n", cap(array))

	// slice
	// We dont specify the amount of elemnts that we will get
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice)
	fmt.Printf("Length of the slice is %d\n", len(slice))
	fmt.Printf("Max capacity of the slice is %d\n", cap(slice))

	// slice method
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// slicing[inclusive: exclusive]
	// first number is inclusive and the second one is exclusive

	// slicing method append
	slice = append(slice, 123)
	fmt.Println("new slice: ", slice)

	// appeng new list
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println("New slicing with lists... ", slice)
}
