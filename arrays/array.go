package main


import "fmt"

func main() {
	var a [5]int
	fmt.Println("Empty 5 element array ", a)

	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])
	/* As in Python the built in function len gives the length of the array */
	fmt.Println("Len of a array: ", len(a))

	/* Initializing and declaring an array */
	b := [5]int{5,4,3,2,1}
	fmt.Println("dcl: ",  b)

	/* Multi dimensional array */
	var twoD [2][3]int // remember, no semicolons!

	for i:=0; i<2;i++{
		for j:=0; j<3; j++{
			twoD[i][j] = i+j
		}
	}

	/*Smarter print function than Cs */
	fmt.Println("2d: ", twoD)
}
