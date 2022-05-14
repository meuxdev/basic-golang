package mypackage

import "fmt"

// CarPublic Public structure to create car
type CarPublic struct {
	Brand           string
	Year            uint
	Model           string
	privateProperty bool
}

type carPrivate struct {
	brand string
	year  uint
}

func privateSecretFunction() {
	fmt.Println("Im just a private function can not be acess from other packages")
}

// Print Public function that just prints a msg
func Print(text string) {
	fmt.Println(text)
}
