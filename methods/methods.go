package main

import "fmt"
/* Go supports methods defined on structs */



type rect struct {
	width, height int
}
/* The area method has a receiver type of rect */
func (r *rect) area() int{
	return r.width * r.height
}
/* Methods can be defined for either pointer or value receiver types. Here's an example of value receiver */
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

/* Go Automatically handles conversion between values and pointers
	for method calls. You may want to use a pointer receiver type
	to avoid copying on the method calls or to allow the method 
	to mutate the receivign struct
*/


func main() {
	r:= rect{width: 10, height: 5}
	fmt.Println("Area: ", r.area())
	fmt.Println("Perim: ", r.perim())
	rp := &r

	fmt.Println("Area: ", rp.area())
	fmt.Println("Perim: ", rp.perim())
}
