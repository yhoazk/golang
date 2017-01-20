package main

import "fmt"
/* Basic sends and receives on channels are blocking
	however, we can use select with a default clause to
	implement non-blocking sends, receives and even non-
	blocking milti-way selects
*/
func main(){
	messages := make(chan string)
	signals  := make(chan bool)

	/* Heres a non-blocking receive. If a value is available on
	messages then select will take the <-message case with
	that value, if not it will immediately take the default 
	case */
	select{
	case msg := <-messages:
		fmt.Println("received Message: ", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	/* A non-blocking send works similarly */
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("No message sent")
	}
	/* We can use multiple cases above that default clause
	to implement a multi-way-non-blocking select. Here
	we attempt non-blocking receives on both messages
	and signals*/
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("No activity")
	}
}
