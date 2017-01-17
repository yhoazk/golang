package main

/* Channels are pipes that connect goroutines.
	you can send values into channels form one
	goroutine and receive those values into another 
	goroutine */

import "fmt"

func main(){
	/*Create a new channel with make(chan val-type)
	  Channels are typed by the values they convey*/
	messages := make(chan string)

	/*SEnd a value into the channel using channle<- syntax */
	go func(){ messages <- "ping" }()
	/* The <-channel syntax receives a value from the channel
	Here we'll receive the ping message */
	msg:= <-messages
	/* By default sends and receives are blocked until both the 
	sender and the receiver are ready, this property allowed is to
	wait at the end of our program for the ping w/o explicit sync*/
	fmt.Println(msg)
}
