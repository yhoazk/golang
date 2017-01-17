package main

import "fmt"
/* Go routines are a ligthweight thread of execution */

func f(from string){
	for i:=0;i<3; i++{
		fmt.Println(from, ":", i)
	}
}


func main(){
	/* This is run synchronously */
	f("direct")
	/* To invoke the function in a go routine use: go <function>  */

	go f("goroutine")
	go f("1234567")
	go f("abcdefghijklmnopq")
	/* A go routine can also be started as an anonymous function */
	go func(msg string){
		fmt.Println(msg)
	}("going")
	/* As the tasks are runnig async in separate go routines now, we
		have to wait for the excecution to fall here and wait for 
		a key press before exiting the programm 
	*/
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}


