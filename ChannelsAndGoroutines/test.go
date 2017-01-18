package main

import (
	"fmt"
	"C"
	"time"
)

func sleepAndTalk(t time.Duration, msg string, c chan string){
	time.Sleep(t * time.Second)
	//fmt.Printf("%v ", msg)
	c<-msg
}

/*
#include <stdlib.h>



func Random() int{
	return int(C.random())
}

func seed(i int)  {
	C.srandom(C.uint(i))

}
*/

func main(){
	//fmt.Printf("asdasdasdads\n")
	c := make(chan string)
	go sleepAndTalk(0, "Hello ", c)
	go sleepAndTalk(1, "Gophers! ", c )
	go sleepAndTalk(2, "What's ", c)
	go sleepAndTalk(3, "up?", c)
	fmt.Println("Done")
	for i:=0; i < 4; i++{
		fmt.Printf("%v", <-c)
	}
	//seed(45)
	//fmt.Printf(strconv.Itoa(Random()))
}
