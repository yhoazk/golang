package main

import "fmt"
/* The signature function shows that vals returns two integers  */
func vals() (int, int){
	return 3,7
}



func main(){
	a,b := vals()
	fmt.Println(a)
	fmt.Println(b)
	/* If only one paramenter is of interest we can use the _ variable as in python */
	_, c := vals()
	fmt.Println(c)
}
