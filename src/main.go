package main

import (
	"fmt"
)

func say(text string, c chan<- string) {
	c <- text
}

func msg(text string, c chan<- string) {
	c <- text
}

func main() {
	// type of val passing to the channel, amount of go routines the channel will receive
	c := make(chan string, 1)
	fmt.Println("Hello")
	go say("Bay", c)

	channel2 := make(chan string, 2)

	channel2 <- "Msg 1"
	channel2 <- "Msg 2"

	// len() -> gives the len of the channel
	// cap() -> capacity of the channel, how many you can add
	fmt.Println(len(channel2), cap(channel2))
	fmt.Println(<-c)

	// Range and close
	// channel2 <- "Msg 3"
	close(channel2)

	for msg := range channel2 {
		fmt.Println(msg)
	}

	email := make(chan string)
	email2 := make(chan string)
	go msg("alex@gmail.com", email)
	go msg("freddy@gmail.com", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email:
			fmt.Println("Email received from", m1)
		case m2 := <-email2:
			fmt.Println("Email received from", m2)
		}
	}

}
