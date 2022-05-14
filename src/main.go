package main

import "fmt"

func main() {
	// Defer
	defer fmt.Println("Hello") // last print waits to all to execute
	// defer executes before main is completed.
	// conexion -> defer cierre conexion
	// abrir archivo -> con defer lo cerramos
	// GOD PRACTICE ONLY USE 1 DEFER PER FUNCTION
	fmt.Println("World")

	// continue and break

	for i := 0; i < 10; i++ {
		fmt.Println(i)

		// continue
		if i == 2 {
			fmt.Println("Es 2")
			continue // example: when a error executes but you dont want to break
		}

		// break
		if i == 8 {
			fmt.Println("Break")
			break // when you want to finalize the cycle
		}
	}
}
