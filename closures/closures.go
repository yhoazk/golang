package main

import "fmt"

/* Go supports anonymous functions, which can from closures. */
/* Anonymous functions are useful when you want to define a
   function inline w/o having to name it
*/

/* the function intSeq returns another fnc defined anonymously in 
the body of the intSeq fnc*/

/* The returned function closes over the variable i to form a 
closure*/

func intSeq() func() int{
	i := 0
	return func() int{
		i+=1
		return i
	}
}





func main(){
	/*	we call intSeq, assigning the result (a function) to nextInt.
	This function value captures its own i value, which will be updated 
	each time we call nextInt.*/
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	/* To confirm that the state is unique to that particular function,
	create and test a new one. */

	fmt.Println("New closure")
	newInts := intSeq()
	fmt.Println(newInts())

}

