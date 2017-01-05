package main

import 	"fmt"
import 	"time"

func main() {
	i := 2
	fmt.Println("Write ", i , " as ")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	/* 	Commas can be used to separate multiple expressions
		in the same case statement.
	*/
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	default:
	fmt.Println("It's Weekday :( ")
	}

	/*	Switch without an expression is an alternative way to
		express if/else logic. Also the case expression can be non
		constant
	*/
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noor")
	}


	/*	A type switch compares types instead of values.
		It can be used to discover the type of an interface value
	*/
	whatAmI := func (i interface{}){
		switch t := i.(type){
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("I'm not bool nor int::- %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("string")



}
