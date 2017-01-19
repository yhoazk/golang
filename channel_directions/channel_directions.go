package main

import "fmt"
/* When using channels as function parameters, its possible
to specify if a channel ismenat to only receive or send values*/
/* This ping fnc only accepts a channel for sendig
of a Rx channels is passed as an argument, it would be a 
compile time error */
func ping(pings chan<-string, msg string){
	pings<-msg
}
/* The pong function accepts one channel for Rx and a 
second for sends*/
func pong(pings <-chan string, pongs chan<-string){
	msg:= <-pings
	pongs <-msg
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}
