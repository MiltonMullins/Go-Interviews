package main

import (
	"fmt"
	"time"
)

/*
Channels are the pipes that connect concurrent goroutines.
You can send values into channels from one goroutine and receive those values into another goroutine.
*/
func channelsFun() {

	// Create a new channel with make(chan val-type). Channels are typed by the values they convey.
	messages := make(chan string)
	/*
		Send a value into a channel using the channel <- syntax.
		Here we send "ping" to the messages channel we made above, from a new goroutine.
	*/
	go func() { messages <- "ping" }()

	/*
		The <-channel syntax receives a value from the channel.
		Here we’ll receive the "ping" message we sent above and print it out.
	*/
	msg := <-messages //Block until we receive a notification from the goroutine.
	fmt.Println(msg)
	/*
		By default sends and receives block until both the sender and receiver are ready.
		This property allowed us to wait at the end of our program for the "ping" message
		without having to use any other synchronization.
	*/
	
	blockingFunTest()
}

func blockingFunTest() {
	done := make(chan bool)

	go func() {
		println("Working...")
		time.Sleep(5 * time.Second)
		println("DONE!")

		done <- false
	}()

	<-done //Block until we receive a notification from the worker on the channel.
	fmt.Println("Test")
}
