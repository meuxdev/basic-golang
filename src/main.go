package main

import "fmt"

func say(text string, c chan string) {
	c <- text
}

func main() {
	// type of val passing to the channel, amount of go routines the channel will receive
	c := make(chan string, 1)
	fmt.Println("Hello")
	go say("Bay", c)
}
