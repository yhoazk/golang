package main

import "fmt"

/* Go supports anonymous functions, which can from closures. */
/* Anonymous functions are useful when you want to define a
   function inline w/o having to name it
*/

/* Closures:
		Los closures son funciones que manejan variables independientes. En otras
		palabras, la funcion definida en el closure "recuerda" el entorno donde
		fue creada.
		https://developer.mozilla.org/es/docs/Web/JavaScript/Closures

		Eso ofrece la posibilidad de emular funciones privadas, es decir,
		que solo pueden ser llamados por otros metodos en la misma clase.
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


/* Emulating private methods */

func counter() func() {
	var privateCounter int;
	var m map
	func mod(n int)  {
		privateCounter += n
	}
	return func() int{

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
