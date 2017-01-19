package main

import "fmt"
import "time"
/* Go's select lets you to wait on multiple channel operations
combining goroutines and channels with select*/

func main(){
	/* we will select across two channels*/
	c1 := make(chan string)
	c2 := make(chan string)
	/* Each channel will receive a value after some time
	to simulate blocking RPC operations in concurrent goroutines */
	go func(){
		time.Sleep(time.Second * 1)
		c1<-"one"
	}()

	go func(){
		time.Sleep(time.Second * 2)
		c2<-"two"
	}()
	/* we 'll use select to await both values simultaneously
	printing each as it arrives*/
	for i:=0; i<2; i++{
		select{
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)

		}
	}
}

/* We receive the values one and two as expected
Note that the total execution time is only 2 seconds
Since both goroutines execute concurrently */

