package main

import "fmt"

func Janis(ch chan string) {
	// This line blocks until a message is sent to the channel
	msg := <-ch

	fmt.Println("Jimi said:", msg)

	// This line blocks until the channel is read from
	ch <- "Hello, Jimi!"
}

func main() {
	// make a new channel of type string
	// and assign it to the phone variable
	phone := make(chan string)

	// Close the channel to signal that no more messages will be sent/received
	defer close(phone)

	// Launch the Janis function as a goroutine
	// This will run concurrently with the main function
	go Janis(phone)

	// This line blocks until there is a line of code ready to read form the channel.
	phone <- "Hello, Janis!"

	// This line blocks until there is a line fo code ready to send a message down the channel
	msg := <-phone
	fmt.Println("Janis said:", msg)
}
